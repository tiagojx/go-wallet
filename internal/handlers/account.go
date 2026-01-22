package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tiagojx/go-wallet/internal/account"
)

type AccountHandler struct {
	repo *account.Repository
}

func NewAccountHandler(repo *account.Repository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// verifica método POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// cria nova struct Account
	acc := new(account.Account)
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		http.Error(w, "error creating account", http.StatusBadRequest)
		return
	}

	// registrando a conta no banco de dados...
	err = h.repo.Save(acc)
	if err != nil {
		http.Error(w, "error creating account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}
