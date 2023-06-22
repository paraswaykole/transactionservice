package transactions

import (
	"strconv"

	"github.com/paraswaykole/transactionservice/internal/db"
)

func CreateTransaction(transaction *Transaction) error {
	err := db.Get().Create(transaction).Error
	return err
}

func GetTransactionByID(id int64) (*Transaction, error) {
	var transaction Transaction
	err := db.Get().Where(&Transaction{ID: id}).Preload("Type").First(&transaction).Error
	return &transaction, err
}

func GetType(name string) (*Type, error) {
	var ttype Type
	err := db.Get().Where(Type{Name: name}).Attrs(Type{Name: name}).FirstOrCreate(&ttype).Error
	return &ttype, err
}

func GetTransactionIDsByTypeID(typeID int) (*[]int64, error) {
	var transactionIDs []int64
	err := db.Get().Table("transactions").Select("id").Where(&Transaction{TypeID: typeID}).Find(&transactionIDs).Error
	return &transactionIDs, err
}

func GetTransactionsSum(tid int64) (*float64, error) {
	var sumAmount *float64
	err := db.Get().Raw("select sum(amount) from transactions where path ~ ?", "*."+strconv.FormatInt(tid, 10)+".*").Scan(&sumAmount).Error
	return sumAmount, err
}
