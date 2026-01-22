package api

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiagojx/go-wallet/internal/handlers"
	"github.com/tiagojx/go-wallet/internal/middleware"
)

type Server struct {
	Router             *mux.Router
	TransactionHandler *handlers.TransactionHandler
	AccountHandler     *handlers.AccountHandler
	Logger             *slog.Logger
}

func NewServer(th *handlers.TransactionHandler, ah *handlers.AccountHandler, logger *slog.Logger) *Server {
	s := &Server{
		Router:             mux.NewRouter(),
		TransactionHandler: th,
		AccountHandler:     ah,
		Logger:             logger,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.Router.Use(middleware.Logging)

	// Endpoints
	s.Router.HandleFunc("/accounts", s.AccountHandler.CreateAccount).Methods("POST")
	s.Router.HandleFunc("/transactions", s.TransactionHandler.CreateTransaction).Methods("POST")

}

func (s *Server) Run(port string) error {
	s.Logger.Info("Running local server on http://localhost:", "port", port)
	return http.ListenAndServe(":"+port, s.Router)
}
