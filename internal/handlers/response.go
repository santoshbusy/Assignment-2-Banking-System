package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type apiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *apiError   `json:"error,omitempty"`
}

func respondSuccess(c *gin.Context, status int, data interface{}) {
	if status == 0 {
		status = http.StatusOK
	}
	c.JSON(status, apiResponse{
		Success: true,
		Data:    data,
	})
}

func respondError(c *gin.Context, status int, code, message string) {
	if status == 0 {
		status = http.StatusInternalServerError
	}
	c.JSON(status, apiResponse{
		Success: false,
		Error: &apiError{
			Code:    code,
			Message: message,
		},
	})
}
