package standardresponse

import (
	"database/sql"
	"errors"
	"net/http"
	businesserror "poc/internal/shared/business_error"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c echo.Context, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusNotFound, StandardResponse[any]{
			Code:    string(businesserror.NotFound),
			Message: err.Error(),
		})
	}

	if businessError, ok := err.(*businesserror.BusinessError); ok {
		return c.JSON(http.StatusBadRequest, StandardResponse[any]{
			Code:    string(businessError.Code()),
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusInternalServerError, StandardResponse[any]{
		Code:    string(businesserror.InternalError),
		Message: err.Error(),
	})
}
