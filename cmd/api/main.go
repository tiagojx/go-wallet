package main

import (
	"fmt"
	"log"

	"github.com/tiagojx/go-wallet/internal/account"
	"github.com/tiagojx/go-wallet/internal/database"
	"github.com/tiagojx/go-wallet/internal/transaction"
)

func main() {
	db, err := database.NewConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	fmt.Println("----------------------------------")
	fmt.Println("Go is now connected to PostgreSQL!")
	fmt.Println("----------------------------------")

	/*
	 * APENAS PARA FINS DE TESTES
	 */

	// inicializando repositórios...
	accRepo := account.NewRepository(db)
	txRepo := transaction.NewRepository(db)

	// Criando contas...
	newAcc := account.NewAccount("Maria")

	if err = accRepo.Save(newAcc); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New account successfully created!\nID: %d | Name: %s | Balance: %d\n",
		newAcc.ID,
		newAcc.Name,
		newAcc.Balance)

	// Processando transações...
	newTx := transaction.NewTransaction(1, 2, 1000)

	if err = txRepo.Create(newTx); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your transaction was successfully created!\nSender: %d | Destination: %d | Amount: %d\n",
		newTx.AccountIdFrom,
		newTx.AccountIdTo,
		newTx.Amount)
}
