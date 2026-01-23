package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tiagojx/go-wallet/internal/transaction"
	"github.com/tiagojx/go-wallet/internal/usecase"
)

type TransactionHandler struct {
	txUsecase *usecase.TransactionUseCase
}

func NewTransactionHandler(txUsecase *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		txUsecase: txUsecase,
	}
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
	if err = h.txUsecase.Execute(tx); err != nil {
		http.Error(w, "error processing transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tx)
}
