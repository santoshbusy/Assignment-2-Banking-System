package models

import "time"

type Bank struct {
	BankID    uint      `gorm:"column:bank_id;primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Branches []Branch `gorm:"foreignKey:BankID"`
	Loans    []Loan   `gorm:"foreignKey:BankID"`
}

func (Bank) TableName() string { return "banks" }
