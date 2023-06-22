package main

import (
	"github.com/paraswaykole/transactionservice/internal/config"
	"github.com/paraswaykole/transactionservice/internal/db"
	"github.com/paraswaykole/transactionservice/internal/server"
	"github.com/paraswaykole/transactionservice/internal/transactions"
)

func main() {
	config.Init()
	db.Init()
	db.Get().AutoMigrate(&transactions.Transaction{}, &transactions.Type{})
	server.Start()
}
