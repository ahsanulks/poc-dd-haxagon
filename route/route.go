package route

import (
	"poc/internal/onboarding"
	"poc/internal/order"

	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo) {
	onboarding.NewHttpRoute(e)
	order.NewHttpRoute(e)
}
