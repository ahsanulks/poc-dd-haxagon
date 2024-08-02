package handler

import (
	"poc/internal/order/adapter"
	"poc/internal/order/usecase"
	standardresponse "poc/internal/shared/standard_response"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type OrderHttpHandler struct {
	orderUsecase adapter.OrderUsecase
}

func NewOrderHttpHandler(orderUsecase adapter.OrderUsecase) *OrderHttpHandler {
	return &OrderHttpHandler{orderUsecase: orderUsecase}
}

func (h *OrderHttpHandler) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	var productRequests []*usecase.ProductRequest
	copier.Copy(&productRequests, &req.ProductRequest)

	order, err := h.orderUsecase.CreateOrder(c.Request().Context(), &usecase.CreateOrderRequest{
		UserID:         c.Get("user_id").(int),
		ProductRequest: productRequests,
		AddressID:      req.AddressID,
	})
	if err != nil {
		return standardresponse.NewErrorResponse(c, err)
	}
	return c.JSON(200, map[string]any{
		"id": order.ID,
	})
}

type CreateOrderRequest struct {
	ProductRequest []*ProductRequest `json:"products"`
	AddressID      int               `json:"address_id"`
}

type ProductRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
