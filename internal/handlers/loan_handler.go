package handlers

import (
	"banking-system/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	service services.LoanService
}

func NewLoanHandler(service services.LoanService) *LoanHandler {
	return &LoanHandler{service: service}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var request struct {
		CustomerID uint    `json:"customer_id"`
		BankID     uint    `json:"bank_id"`
		BranchID   uint    `json:"branch_id"`
		Principal  float64 `json:"principal"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.CustomerID == 0 || request.BankID == 0 || request.BranchID == 0 || request.Principal <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_FIELDS", "customer_id, bank_id, branch_id and positive principal are required")
		return
	}

	loan, err := h.service.CreateLoan(
		request.CustomerID,
		request.BankID,
		request.BranchID,
		request.Principal,
	)

	if err != nil {
		respondError(c, http.StatusBadRequest, "CREATE_LOAN_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusCreated, gin.H{"loan": loan})
}

func (h *LoanHandler) RepayLoan(c *gin.Context) {
	loanIDParam := c.Param("id")
	loanID, err := strconv.Atoi(loanIDParam)
	if err != nil || loanID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_LOAN_ID", "loan id must be a positive integer")
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.Amount <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_AMOUNT", "amount must be greater than zero")
		return
	}

	loan, err := h.service.RepayLoan(uint(loanID), request.Amount)
	if err != nil {
		respondError(c, http.StatusBadRequest, "REPAY_LOAN_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, gin.H{
		"message": "loan repayment successful",
		"loan":    loan,
	})
}

func (h *LoanHandler) GetLoanDetails(c *gin.Context) {
	loanIDParam := c.Param("id")
	loanID, err := strconv.Atoi(loanIDParam)
	if err != nil || loanID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_LOAN_ID", "loan id must be a positive integer")
		return
	}

	loan, interest, err := h.service.GetLoanDetails(uint(loanID))
	if err != nil {
		respondError(c, http.StatusNotFound, "LOAN_NOT_FOUND", "loan not found")
		return
	}

	respondSuccess(c, http.StatusOK, gin.H{
		"loan":          loan,
		"interest_year": interest,
	})
}
