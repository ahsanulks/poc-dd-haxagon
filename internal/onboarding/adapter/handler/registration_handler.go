package handler

import (
	"poc/internal/onboarding/adapter"
	"poc/internal/onboarding/usecase"
	standardresponse "poc/internal/shared/standard_response"

	"github.com/labstack/echo/v4"
)

type RegistrationHttpHandler struct {
	registrationUsecase adapter.RegistrationUsecase
}

func NewRegistrationHttpHandler(registrationUsecase adapter.RegistrationUsecase) *RegistrationHttpHandler {
	return &RegistrationHttpHandler{
		registrationUsecase: registrationUsecase,
	}
}

type RegistrationBody struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (h *RegistrationHttpHandler) Register(c echo.Context) error {
	var req RegistrationBody
	if err := c.Bind(&req); err != nil {
		return err
	}

	token, err := h.registrationUsecase.Register(c.Request().Context(), &usecase.RegistrationRequest{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		return standardresponse.NewErrorResponse(c, err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}
