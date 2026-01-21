package transaction

import "time"

type Transaction struct {
	ID            int       `json:"id"`
	AccountIdFrom int       `json:"account_id_from"`
	AccountIdTo   int       `json:"account_id_to"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewTransaction(accountIdFrom int, accountIdTo int, amount int64) *Transaction {
	return &Transaction{
		AccountIdFrom: accountIdFrom,
		AccountIdTo:   accountIdTo,
		Amount:        amount,
	}
}
