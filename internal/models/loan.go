package models

import "time"

type Loan struct {
	LoanID          uint      `gorm:"column:loan_id;primaryKey"`
	CustomerID      uint      `gorm:"column:customer_id;not null"`
	BankID          uint      `gorm:"column:bank_id;not null"`
	BranchID        uint      `gorm:"column:branch_id;not null"`
	Principal       float64   `gorm:"column:principal;not null"`
	InterestRate    float64   `gorm:"column:interest_rate;not null"`
	RemainingAmount float64   `gorm:"column:remaining_amount;not null"`
	CreatedAt       time.Time `gorm:"column:created_at"`

	Customer Customer      `gorm:"foreignKey:CustomerID"`
	Bank     Bank          `gorm:"foreignKey:BankID"`
	Branch   Branch        `gorm:"foreignKey:BranchID"`
	Payments []LoanPayment `gorm:"foreignKey:LoanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Loan) TableName() string { return "loans" }
