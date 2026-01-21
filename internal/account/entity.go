package account

import "time"

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(name string) *Account {
	return &Account{
		Name:    name,
		Balance: 0,
	}
}
