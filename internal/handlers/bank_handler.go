package handlers

import (
	"banking-system/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	service services.BankService
}

func NewBankHandler(service services.BankService) *BankHandler {
	return &BankHandler{service: service}
}

func (h *BankHandler) CreateBank(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.Name == "" {
		respondError(c, http.StatusBadRequest, "INVALID_NAME", "name is required")
		return
	}

	bank, err := h.service.CreateBank(request.Name)
	if err != nil {
		respondError(c, http.StatusBadRequest, "CREATE_BANK_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusCreated, gin.H{"bank": bank})
}

func (h *BankHandler) GetBanks(c *gin.Context) {
	banks, err := h.service.GetBanks()
	if err != nil {
		respondError(c, http.StatusInternalServerError, "FETCH_BANKS_FAILED", "failed to fetch banks")
		return
	}

	respondSuccess(c, http.StatusOK, banks)
}
