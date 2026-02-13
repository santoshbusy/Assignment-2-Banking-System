package models

import "time"

type Branch struct {
	BranchID  uint      `gorm:"column:branch_id;primaryKey"`
	BankID    uint      `gorm:"column:bank_id;not null"`
	Name      string    `gorm:"column:name;not null"`
	Address   string    `gorm:"column:address;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Bank     Bank      `gorm:"foreignKey:BankID;references:BankID"`
	Accounts []Account `gorm:"foreignKey:BranchID"`
	Loans    []Loan    `gorm:"foreignKey:BranchID"`
}

func (Branch) TableName() string { return "branches" }
