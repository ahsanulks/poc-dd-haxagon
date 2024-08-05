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

// AddAddress
// @Summary Add address
// @Description Add new user address
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Router /users/addresses [post]
// @Param body body AddAddressBody true "address data"
// @Security BearerAuth
// @Success 201 {object} standardresponse.StandardResponse[CreateAddressResponse]
// @Failure 400 {object} standardresponse.StandardResponse[any]
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

	return standardresponse.NewSuccessResponse(c, 201, CreateAddressResponse{
		ID: address.ID,
	})
}

// GetAddresses
// @Summary Get addresses
// @Description Get user addresses
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Router /users/addresses [get]
// @Security BearerAuth
// @Success 200 {object} standardresponse.StandardResponse[[]AddressResponse]
// @Failure 400 {object} standardresponse.StandardResponse[any]
func (h *AddressHttpHandler) GetAddress(c echo.Context) error {
	userID := c.Get("user_id").(int)
	addresses, err := h.addressManagement.GetAddress(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	var addressResponses []*AddressResponse
	copier.Copy(&addressResponses, addresses)

	return standardresponse.NewSuccessResponse(c, 200, addressResponses)
}

type AddressResponse struct {
	ID        int     `json:"id"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateAddressResponse struct {
	ID int `json:"id"`
}
