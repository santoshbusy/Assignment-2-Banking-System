package handlers

import (
	"banking-system/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service services.CustomerService
}

func NewCustomerHandler(service services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		respondError(c, http.StatusBadRequest, "INVALID_BODY", "invalid request body")
		return
	}

	if request.Name == "" || request.Email == "" {
		respondError(c, http.StatusBadRequest, "INVALID_FIELDS", "name and email are required")
		return
	}

	customer, err := h.service.CreateCustomer(request.Name, request.Email)
	if err != nil {
		respondError(c, http.StatusBadRequest, "CREATE_CUSTOMER_FAILED", err.Error())
		return
	}

	respondSuccess(c, http.StatusCreated, gin.H{"customer": customer})
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, "INVALID_CUSTOMER_ID", "customer id must be a positive integer")
		return
	}

	customer, err := h.service.GetCustomer(uint(id))
	if err != nil {
		respondError(c, http.StatusNotFound, "CUSTOMER_NOT_FOUND", "customer not found")
		return
	}

	respondSuccess(c, http.StatusOK, customer)
}
