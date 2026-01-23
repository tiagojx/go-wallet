package main

import (
	"github.com/tiagojx/go-wallet/internal/account"
	"github.com/tiagojx/go-wallet/internal/api"
	"github.com/tiagojx/go-wallet/internal/database"
	"github.com/tiagojx/go-wallet/internal/event"
	"github.com/tiagojx/go-wallet/internal/handlers"
	"github.com/tiagojx/go-wallet/internal/middleware"
	"github.com/tiagojx/go-wallet/internal/transaction"
	"github.com/tiagojx/go-wallet/internal/usecase"
)

func main() {
	/*
	 * Iniciando o middleware...
	 */

	logger := middleware.InitLogger()
	logger.Info("starting go-wallet API...")

	/*
	 * Conectando ao banco de dados...
	 */

	db, err := database.NewConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger.Info("Go is now connected to PostgreSQL!")

	/*
	 * Inicializando repositórios...
	 */

	accRepo := account.NewRepository(db)
	txRepo := transaction.NewRepository(db)

	/*
	 * Registrando eventos...
	 */

	rabbitConnStr := event.NewConnection()
	queueName := "transaction_queue"

	prod, err := event.NewProducer(rabbitConnStr, queueName)
	if err != nil {
		panic(err)
	}
	defer prod.Close()

	consumer, err := event.NewConsumer(rabbitConnStr)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	/*
	 * Registrando UseCases
	 */

	txUsecase := usecase.CreateTransactionUseCase(txRepo, prod)

	/*
	 * Registrando handlers...
	 */

	accHandler := handlers.NewAccountHandler(accRepo)
	txHandler := handlers.NewTransactionHandler(txUsecase)

	/*
	 * Leitura das mensagens do RabbitMQ
	 */

	go func() {
		if err := consumer.Start(queueName); err != nil {
			logger.Error("error in rabbitmq consumer", "error", err)
		}
	}()

	/*
	 * Server
	 */
	server := api.NewServer(txHandler, accHandler, logger)
	if err = server.Run("8080"); err != nil {
		panic(err)
	}
}
