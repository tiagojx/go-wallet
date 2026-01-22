package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tiagojx/go-wallet/internal/transaction"
)

type TransactionHandler struct {
	repo *transaction.Repository
}

func NewTransactionHandler(repo *transaction.Repository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// verifica se o método usado é POST.
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// decodifica o json do corpo da requsição para uma scrutct do tipo Transaction.
	tx := new(transaction.Transaction)
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "error decoding json", http.StatusBadRequest)
		return
	}

	// criando transação...
	if err = h.repo.Create(tx); err != nil {
		http.Error(w, "error processing transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tx)
}
