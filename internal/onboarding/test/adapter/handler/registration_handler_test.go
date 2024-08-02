package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"poc/internal/onboarding/adapter/handler"
	"poc/internal/onboarding/test/mock"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegistrationHttpHandler_Register(t *testing.T) {
	tests := []struct {
		name               string
		params             map[string]any
		expectedStatusCode int
		mcFunc             func(mock *mock.MockRegistrationUsecase)
	}{
		{
			name: "success",
			params: map[string]any{
				"name":         "test",
				"phone_number": "08123112313",
			},
			expectedStatusCode: http.StatusCreated,
			mcFunc: func(mock *mock.MockRegistrationUsecase) {
				mock.EXPECT().Register(gomock.Any(), gomock.Any()).Return("token", nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			regUsecase := mock.NewMockRegistrationUsecase(ctrl)
			tt.mcFunc(regUsecase)

			h := handler.NewRegistrationHttpHandler(regUsecase)

			assert := assert.New(t)

			jsonBody, err := json.Marshal(tt.params)
			assert.NoError(err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err = h.Register(c)

			assert.NoError(err)
			assert.Equal(tt.expectedStatusCode, rec.Result().StatusCode)
		})
	}
}
