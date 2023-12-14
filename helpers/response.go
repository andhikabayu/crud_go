package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func SuccessJSON(c *gin.Context, message string, data interface{}) {
	response := SuccessResponse{
		Status:  1,
		Message: message,
		Error:   false,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func ErrorJSON(c *gin.Context, message string) {
	response := ErrorResponse{
		Status:  0,
		Message: message,
		Error:   true,
	}
	c.JSON(http.StatusBadRequest, response)
}
