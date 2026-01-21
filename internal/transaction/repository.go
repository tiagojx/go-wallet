package transaction

import (
	"database/sql"
	"fmt"
)

type RepositoryInterface interface {
	Create(transaction *Transaction) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(t *Transaction) error {
	// iniciando SQL Transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error when starting the transaction: %v", err)
	}

	defer tx.Rollback() // programar a reversão dos dados em caso de falha.

	// escrevendo o novo registro de transação no banco de dados.
	stmt := "INSERT INTO transactions (account_id_from, account_id_to, amount) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRow(stmt, t.AccountIdFrom, t.AccountIdTo, t.Amount).Scan(&t.ID)
	if err != nil {
		return fmt.Errorf("error when creating the transaction log: %v", err)
	}

	// atualizando a conta do remetente.
	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", t.Amount, t.AccountIdFrom)
	if err != nil {
		return fmt.Errorf("error when debiting the sender account: %v", err)
	}

	// atualizando a conta do destinatário.
	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", t.Amount, t.AccountIdTo)
	if err != nil {
		return fmt.Errorf("error when crediting destination account: %v", err)
	}

	// commitando a atualização permanente do DB.
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error when commiting the transaction to the database: %v", err)
	}

	return nil
}
