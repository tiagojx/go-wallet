package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tiagojx/go-wallet/internal/account"
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

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.HandleFunc("/accounts", accHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/transactions", txHandler.CreateTransaction).Methods("POST")

	logger.Info("Running local server on http://localhost:8080/")
	if err = http.ListenAndServe(":8080", r); err != nil {
		logger.Error("error when starting server", "error", err)
		os.Exit(1)
	}
}
