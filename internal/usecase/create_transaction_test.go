package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/tiagojx/go-wallet/internal/transaction"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(tx *transaction.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}
