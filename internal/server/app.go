package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paraswaykole/transactionservice/internal/transactions"
)

func Start() {
	app := fiber.New()

	serviceGroup := app.Group("/transactionservice")
	{
		transactionsGroup := serviceGroup.Group("transaction")
		{
			transactionsGroup.Put(":tid", transactions.PutController)
			transactionsGroup.Get(":tid", transactions.GetController)
		}
		typesGroup := serviceGroup.Group("types")
		{
			typesGroup.Get(":type", transactions.GetTypeController)
		}
		sumsGroup := serviceGroup.Group("sum")
		{
			sumsGroup.Get(":tid", transactions.GetSumController)
		}
	}

	app.Listen(":3000")
}
