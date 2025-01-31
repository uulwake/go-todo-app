package libs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DefaultInternalServerError(ctx echo.Context, err error) CustomError {
	return CustomError{
		HTTPCode: http.StatusInternalServerError,
		Message: err.Error(),
	}
}

type CustomError struct {
	HTTPCode int 
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}