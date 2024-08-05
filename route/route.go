package route

import (
	"poc/internal/onboarding"
	"poc/internal/order"

	_ "poc/docs"

	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Bearer {token}". e.g Bearer thisIsMySecurityToken

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func NewRoute(e *echo.Echo) {
	e.GET("/swagger/*", swagger.WrapHandler)
	apiV1 := e.Group("/api/v1")
	onboarding.NewHttpRoute(apiV1)
	order.NewHttpRoute(apiV1)
}
