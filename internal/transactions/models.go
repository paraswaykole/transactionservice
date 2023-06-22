package transactions

import "strconv"

type Transaction struct {
	ID       int64   `gorm:"primaryKey"`
	Amount   float64 `gorm:"not null"`
	TypeID   int     `gorm:"not null"`
	Path     string  `gorm:"type:ltree;not null"`
	ParentID *int64

	Parent *Transaction `gorm:"foreignKey:ParentID"`
	Type   Type         `gorm:"foreignKey:TypeID"`
}

type Type struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}

func NewTransaction(id int64, amount float64, typeID int) *Transaction {
	return &Transaction{
		ID:     id,
		Amount: amount,
		TypeID: typeID,
		Path:   strconv.FormatInt(id, 10),
	}
}

func (t Transaction) NewChildTransaction(id int64, amount float64, typeID int) *Transaction {
	return &Transaction{
		ID:       id,
		Amount:   amount,
		TypeID:   typeID,
		ParentID: &t.ID,
		Path:     t.Path + "." + strconv.FormatInt(id, 10),
	}
}

func (t Transaction) ToView() map[string]interface{} {
	return map[string]interface{}{
		"amount":    t.Amount,
		"type":      t.Type.Name,
		"parent_id": t.ParentID,
	}
}
