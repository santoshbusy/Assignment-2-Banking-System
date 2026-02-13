package models

import "time"

type Customer struct {
	CustomerID uint      `gorm:"column:customer_id;primaryKey"`
	Name       string    `gorm:"column:name;not null"`
	Email      string    `gorm:"column:email;unique;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`

	Accounts []Account `gorm:"foreignKey:CustomerID"`
	Loans    []Loan    `gorm:"foreignKey:CustomerID"`
}

func (Customer) TableName() string { return "customers" }
