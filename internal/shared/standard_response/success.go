package standardresponse

import (
	"github.com/labstack/echo/v4"
)

func NewSuccessResponse(c echo.Context, httpCode int, data any) error {
	return c.JSON(httpCode, StandardResponse[any]{
		Code:    "S000",
		Message: "success",
		Data:    data,
	})
}
