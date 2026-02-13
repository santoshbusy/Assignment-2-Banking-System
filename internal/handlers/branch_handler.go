package handlers

import (
	"banking-system/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchHandler struct {
	service services.BranchService
}

func NewBranchHandler(service services.BranchService) *BranchHandler {
	return &BranchHandler{service: service}
}

func (h *BranchHandler) CreateBranch(c *gin.Context) {
	var request struct {
		BankID  uint   `json:"bank_id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.BankID == 0 || request.Name == "" || request.Address == "" {
		respondError(c, http.StatusBadRequest, "INVALID_FIELDS", "bank_id, name and address are required")
		return
	}

	branch, err := h.service.CreateBranch(request.BankID, request.Name, request.Address)
	if err != nil {
		respondError(c, http.StatusBadRequest, "CREATE_BRANCH_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusCreated, gin.H{"branch": branch})
}

func (h *BranchHandler) GetBranches(c *gin.Context) {
	bankIDStr := c.Query("bank_id")
	if bankIDStr == "" {
		respondError(c, http.StatusBadRequest, "MISSING_BANK_ID", "bank_id query parameter is required")
		return
	}

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil || bankID <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_BANK_ID", "bank_id must be a positive integer")
		return
	}

	branches, err := h.service.GetBranches(uint(bankID))
	if err != nil {
		respondError(c, http.StatusInternalServerError, "FETCH_BRANCHES_FAILED", "failed to fetch branches")
		return
	}

	respondSuccess(c, http.StatusOK, branches)
}
