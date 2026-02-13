package models

import "time"

type LoanPayment struct {
	PaymentID  uint      `gorm:"column:payment_id;primaryKey"`
	LoanID     uint      `gorm:"column:loan_id;not null"`
	AmountPaid float64   `gorm:"column:amount_paid;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (LoanPayment) TableName() string { return "loan_payments" }
