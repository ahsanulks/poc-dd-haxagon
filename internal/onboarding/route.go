package onboarding

import (
	"os"
	"poc/internal/onboarding/adapter/handler"
	"poc/internal/onboarding/adapter/repository"
	"poc/internal/onboarding/usecase"
	"poc/internal/shared/db"
	"poc/internal/shared/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewHttpRoute(e *echo.Echo) {
	dbPool := db.NewPostgresqlConn()
	handlerProvider := NewHandlerProvider(dbPool)

	userGroup := e.Group("users")
	userGroup.POST("/register", handlerProvider.registrationHandler.Register)

	userGroup.POST("/addresses", handlerProvider.addressHandler.AddAddress, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
	userGroup.GET("/addresses", handlerProvider.addressHandler.GetAddress, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
}

type HandlerProvider struct {
	registrationHandler *handler.RegistrationHttpHandler
	addressHandler      *handler.AddressHttpHandler
}

func NewHandlerProvider(dbPool *pgxpool.Pool) *HandlerProvider {
	userPostgresql := repository.NewUserPostgresql(dbPool)

	tokenGenerator := usecase.NewJWTToken(os.Getenv("JWT_SECRET"), os.Getenv("JWT_PUBLIC"))

	registrationUsecase := usecase.NewRegistrationUsecase(userPostgresql, tokenGenerator)
	registrationHandler := handler.NewRegistrationHttpHandler(registrationUsecase)

	addressUsecase := usecase.NewAddressUsecase(userPostgresql, userPostgresql)
	addressHandler := handler.NewAddressHttpHandler(addressUsecase)

	return &HandlerProvider{
		registrationHandler: registrationHandler,
		addressHandler:      addressHandler,
	}
}
