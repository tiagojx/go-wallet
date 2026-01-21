package account

import (
	"database/sql"
	"fmt"
)

type RepositoryInterface interface {
	Save(a *Account) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(a *Account) error {
	stmt := "INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id"

	err := r.db.QueryRow(stmt, a.Name, a.Balance).Scan(&a.ID)
	if err != nil {
		return fmt.Errorf("error creating account: %v", err)
	}

	return nil
}
