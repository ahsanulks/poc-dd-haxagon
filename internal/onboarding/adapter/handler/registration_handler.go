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

// Register new user
// @Summary Register new user
// @Description Register new user
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Router /users/register [post]
// @Param body body RegistrationBody true "User data"
// @Success 201 {object} standardresponse.StandardResponse[RegistrationResponse]
// @Failure 400 {object} standardresponse.StandardResponse[any]
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

	return standardresponse.NewSuccessResponse(c, 201, RegistrationResponse{
		Token: token,
	})
}

type RegistrationResponse struct {
	Token string `json:"token"`
}
