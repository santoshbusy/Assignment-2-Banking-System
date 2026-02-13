package handlers

import (
	"banking-system/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service services.AccountService
}

func NewAccountHandler(service services.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// Deposit Handler
func (h *AccountHandler) Deposit(c *gin.Context) {
	accountIDParam := c.Param("id")
	accountID, err := strconv.Atoi(accountIDParam)
	if err != nil || accountID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_ACCOUNT_ID", "account id must be a positive integer")
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

	account, err := h.service.Deposit(uint(accountID), request.Amount)
	if err != nil {
		respondError(c, http.StatusBadRequest, "DEPOSIT_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, gin.H{
		"message":     "deposit successful",
		"transaction": "DEPOSIT",
		"amount":      request.Amount,
		"account":     account,
	})
}

// Withdraw Handler
func (h *AccountHandler) Withdraw(c *gin.Context) {
	accountIDParam := c.Param("id")
	accountID, err := strconv.Atoi(accountIDParam)
	if err != nil || accountID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_ACCOUNT_ID", "account id must be a positive integer")
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

	account, err := h.service.Withdraw(uint(accountID), request.Amount)
	if err != nil {
		respondError(c, http.StatusBadRequest, "WITHDRAW_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusOK, gin.H{
		"message":     "withdraw successful",
		"transaction": "WITHDRAW",
		"amount":      request.Amount,
		"account":     account,
	})
}

// GET ACCOUNT HANDLER
func (h *AccountHandler) GetAccount(c *gin.Context) {
	accountIDParam := c.Param("id")
	accountID, err := strconv.Atoi(accountIDParam)
	if err != nil || accountID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_ACCOUNT_ID", "account id must be a positive integer")
		return
	}

	account, err := h.service.GetAccount(uint(accountID))
	if err != nil {
		respondError(c, http.StatusNotFound, "ACCOUNT_NOT_FOUND", "account not found")
		return
	}

	respondSuccess(c, http.StatusOK, account)
}

func (h *AccountHandler) OpenAccount(c *gin.Context) {
	var request struct {
		CustomerID uint `json:"customer_id"`
		BankID     uint `json:"bank_id"`
		BranchID   uint `json:"branch_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.CustomerID == 0 || request.BankID == 0 || request.BranchID == 0 {
		respondError(c, http.StatusBadRequest, "INVALID_FIELDS", "customer_id, bank_id and branch_id are required")
		return
	}

	account, err := h.service.OpenAccount(request.CustomerID, request.BankID, request.BranchID)
	if err != nil {
		respondError(c, http.StatusBadRequest, "CREATE_ACCOUNT_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusCreated, gin.H{"account": account})
}

func (h *AccountHandler) Transfer(c *gin.Context) {
	var request struct {
		FromAccountID uint    `json:"from_account_id"`
		ToAccountID   uint    `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.FromAccountID == 0 || request.ToAccountID == 0 || request.Amount <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_FIELDS", "from_account_id, to_account_id and positive amount are required")
		return
	}

	if request.FromAccountID == request.ToAccountID {
		respondError(c, http.StatusBadRequest, "INVALID_ACCOUNTS", "source and destination accounts must differ")
		return
	}

	if err := h.service.Transfer(
		request.FromAccountID,
		request.ToAccountID,
		request.Amount,
	); err != nil {
		respondError(c, http.StatusBadRequest, "TRANSFER_FAILED", err.Error())
		return
	}

	fromAccount, _ := h.service.GetAccount(request.FromAccountID)
	toAccount, _ := h.service.GetAccount(request.ToAccountID)

	respondSuccess(c, http.StatusOK, gin.H{
		"message":      "transfer successful",
		"from_account": fromAccount,
		"to_account":   toAccount,
		"amount":       request.Amount,
	})
}

func (h *AccountHandler) GetTransactions(c *gin.Context) {
	accountIDParam := c.Param("id")
	accountID, err := strconv.Atoi(accountIDParam)
	if err != nil || accountID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_ACCOUNT_ID", "account id must be a positive integer")
		return
	}

	transactions, err := h.service.GetTransactions(uint(accountID))
	if err != nil {
		respondError(c, http.StatusNotFound, "TRANSACTIONS_NOT_FOUND", "transactions not found")
		return
	}

	respondSuccess(c, http.StatusOK, transactions)
}
