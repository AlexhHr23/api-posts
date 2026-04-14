package handlers

import (
	"net/http"

	"github.com/AlexhHr23/gopost-api/server"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func RespondError(c *server.Context, appErr *AppError) {
	c.JSON(appErr.Code, ErrorResponse{
		Error:   http.StatusText(appErr.Code),
		Message: appErr.Message,
		Code:    appErr.Code,
	})
}
