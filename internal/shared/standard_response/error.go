package standardresponse

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c echo.Context, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(404, ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(500, ErrorResponse{
		Message: err.Error(),
	})
}
