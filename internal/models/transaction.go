package models

import "time"

type Transaction struct {
	TransactionID uint      `gorm:"column:transaction_id;primaryKey"`
	AccountID     uint      `gorm:"column:account_id;not null"`
	Type          string    `gorm:"column:type;not null"`
	Amount        float64   `gorm:"column:amount;not null"`
	CreatedAt     time.Time `gorm:"column:created_at"`

	Account Account `gorm:"foreignKey:AccountID"`
}

func (Transaction) TableName() string { return "transactions" }
