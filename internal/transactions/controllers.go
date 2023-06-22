package transactions

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PutController(c *fiber.Ctx) error {
	var body struct {
		Amount   float64 `json:"amount"`
		Type     string  `json:"type"`
		ParentID *int64  `json:"parent_id,omitempty"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  err.Error(),
		})
	}
	tid, err := strconv.ParseInt(c.Params("tid", ""), 10, 64)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  "invalid id",
		})
	}

	ttype, _ := GetType(body.Type)

	var transaction *Transaction
	if body.ParentID == nil {
		transaction = NewTransaction(tid, body.Amount, ttype.ID)
	} else {
		parentTransaction, err := GetTransactionByID(*body.ParentID)
		if err != nil {
			return c.JSON(map[string]interface{}{
				"status": "failed",
				"error":  "invalid parent_id",
			})
		}
		transaction = parentTransaction.NewChildTransaction(tid, body.Amount, ttype.ID)
	}

	err = CreateTransaction(transaction)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  err.Error(),
		})
	}

	return c.JSON(map[string]interface{}{
		"status": "ok",
	})
}

func GetController(c *fiber.Ctx) error {
	tid, err := strconv.ParseInt(c.Params("tid", ""), 10, 64)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  "invalid id",
		})
	}

	transaction, err := GetTransactionByID(tid)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  "transaction not found",
		})
	}

	return c.JSON(transaction.ToView())
}

func GetTypeController(c *fiber.Ctx) error {
	typeName := c.Params("type", "")
	if typeName == "" {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  "empty type",
		})
	}

	ttype, err := GetType(typeName)
	if err != nil {
		return c.JSON([]int64{})
	}

	transactionIDs, err := GetTransactionIDsByTypeID(ttype.ID)
	if err != nil {
		return c.JSON([]int64{})
	}

	return c.JSON(transactionIDs)
}

func GetSumController(c *fiber.Ctx) error {
	tid, err := strconv.ParseInt(c.Params("tid", ""), 10, 64)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
			"error":  "invalid id",
		})
	}

	sum, err := GetTransactionsSum(tid)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"status": "failed",
		})
	}

	return c.JSON(map[string]interface{}{
		"sum": sum,
	})
}
