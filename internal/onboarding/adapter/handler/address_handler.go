package handler

import (
	"poc/internal/onboarding/adapter"
	"poc/internal/onboarding/usecase"
	standardresponse "poc/internal/shared/standard_response"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type AddressHttpHandler struct {
	addressManagement adapter.AddressManagementUsecase
}

func NewAddressHttpHandler(addressManagement adapter.AddressManagementUsecase) *AddressHttpHandler {
	return &AddressHttpHandler{
		addressManagement: addressManagement,
	}
}

type AddAddressBody struct {
	Street    string  `json:"street"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (h *AddressHttpHandler) AddAddress(c echo.Context) error {
	var req AddAddressBody
	if err := c.Bind(&req); err != nil {
		return err
	}

	address, err := h.addressManagement.AddAddress(c.Request().Context(), &usecase.AddAddressRequest{
		UserID:    c.Get("user_id").(int),
		Street:    req.Street,
		City:      req.City,
		ZipCode:   req.ZipCode,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})
	if err != nil {
		return standardresponse.NewErrorResponse(c, err)
	}

	return c.JSON(201, map[string]any{
		"id": address.ID,
	})
}

func (h *AddressHttpHandler) GetAddress(c echo.Context) error {
	userID := c.Get("user_id").(int)
	addresses, err := h.addressManagement.GetAddress(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	var addressResponses []*AddressResponse
	copier.Copy(&addressResponses, addresses)

	return c.JSON(200, map[string]any{
		"addresses": addresses,
	})
}

type AddressResponse struct {
	ID        int     `json:"id"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
