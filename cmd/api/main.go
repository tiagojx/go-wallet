package main

import (
	"fmt"
	"log"
	
	"github.com/tiagojx/go-wallet/internal/account"
	"github.com/tiagojx/go-wallet/internal/database"
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

	repo := account.NewRepository(db)

	newAcc := account.NewAccount("Cléber")

	if err = repo.Save(newAcc); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New account was created successfully!\nID: %d | Name: %s | Balance: %d\n", newAcc.ID, newAcc.Name, newAcc.Balance)
}
