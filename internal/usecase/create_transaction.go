package usecase

import (
	"encoding/json"

	"github.com/tiagojx/go-wallet/internal/event"
	"github.com/tiagojx/go-wallet/internal/transaction"
)

type TransactionUseCase struct {
	txRepo   transaction.RepositoryInterface
	producer event.ProducerInterface
}

func CreateTransactionUseCase(txRepo *transaction.Repository, prod *event.Producer) *TransactionUseCase {
	return &TransactionUseCase{
		txRepo:   txRepo,
		producer: prod,
	}
}

func (txUsecase *TransactionUseCase) Execute(txData *transaction.Transaction) error {
	if err := txUsecase.txRepo.Create(txData); err != nil {
		return err
	}

	txDataMq, err := json.Marshal(txData)
	if err != nil {
		return err
	}

	err = txUsecase.producer.Publish(txDataMq)
	if err != nil {
		return err
	}

	return nil
}
