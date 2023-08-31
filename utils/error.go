package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	StatusCode          int
	ResponseCode        string
	ResponseDescription string
}

// api errors
var (
	ErrDuplicateData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "001",
		ResponseDescription: "Error: Data already exist",
	}
	ErrInsertData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "002",
		ResponseDescription: "Error: Failed inserting data",
	}
	ErrReadData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "003",
		ResponseDescription: "Error: Failed fetching data",
	}
	ErrUpdateData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "004",
		ResponseDescription: "Error: Failed updating data",
	}
	ErrDeleteData = APIError{
		StatusCode:          http.StatusBadRequest,
		ResponseCode:        "005",
		ResponseDescription: "Error: Failed deleting data. Dependency?",
	}
)

func ErrorMessage(c *gin.Context, apiError APIError) {
	c.Abort()
	c.JSON(apiError.StatusCode, gin.H{
		"responseCode":        apiError.ResponseCode,
		"responseDescription": apiError.ResponseDescription,
	})
}
