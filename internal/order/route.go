package order

import (
	"os"
	"poc/internal/order/adapter/handler"
	"poc/internal/order/adapter/repository"
	"poc/internal/order/adapter/repository/order"
	"poc/internal/order/usecase"
	"poc/internal/shared/db"
	"poc/internal/shared/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewHttpRoute(e *echo.Group) {
	dbPool := db.NewPostgresqlConn()
	handlerProvider := NewHandlerProvider(dbPool)

	userGroup := e.Group("/orders")
	userGroup.POST("", handlerProvider.orderHandler.CreateOrder, middleware.JWTMiddleware(os.Getenv("JWT_SECRET")))
}

type HandlerProvider struct {
	orderHandler *handler.OrderHttpHandler
}

func NewHandlerProvider(dbPool *pgxpool.Pool) *HandlerProvider {
	orderPg := order.NewOrderPostgresql(dbPool, order.New(dbPool))
	addressPg := repository.NewAddressPostgresql(dbPool)
	productMem := repository.NewProductInMemory()

	orderUsecase := usecase.NewOrderUsecase(addressPg, productMem, orderPg)
	orderHandler := handler.NewOrderHttpHandler(orderUsecase)

	return &HandlerProvider{
		orderHandler: orderHandler,
	}
}
