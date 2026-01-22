package main

import (
	"os"

	"github.com/tiagojx/go-wallet/internal/account"
	"github.com/tiagojx/go-wallet/internal/api"
	"github.com/tiagojx/go-wallet/internal/database"
	"github.com/tiagojx/go-wallet/internal/handlers"
	"github.com/tiagojx/go-wallet/internal/middleware"
	"github.com/tiagojx/go-wallet/internal/transaction"
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
		logger.Info("error connecting to the database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("Go is now connected to PostgreSQL!")

	/*
	 * Inicializando repositórios...
	 */

	accRepo := account.NewRepository(db)
	txRepo := transaction.NewRepository(db)

	/*
	 * Registrando handlers...
	 */

	accHandler := handlers.NewAccountHandler(accRepo)
	txHandler := handlers.NewTransactionHandler(txRepo)

	/*
	 * Server
	 */
	server := api.NewServer(txHandler, accHandler, logger)
	if err = server.Run("8080"); err != nil {
		logger.Info("error starting server", "error", err)
		os.Exit(1)
	}
}
