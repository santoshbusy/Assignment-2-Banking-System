package models

import "time"

type Account struct {
	AccountID  uint      `gorm:"column:account_id;primaryKey"`
	CustomerID uint      `gorm:"column:customer_id;not null"`
	BankID     uint      `gorm:"column:bank_id;not null"`
	BranchID   uint      `gorm:"column:branch_id;not null"`
	Balance    float64   `gorm:"column:balance;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Type       string    `gorm:"column:type;default:'SAVINGS'"`

	Customer     Customer      `gorm:"foreignKey:CustomerID"`
	Branch       Branch        `gorm:"foreignKey:BranchID"`
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
}

func (Account) TableName() string { return "accounts" }
