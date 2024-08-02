package usecase_test

import (
	"context"
	"poc/internal/onboarding/test/mock"
	"poc/internal/onboarding/usecase"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestRegistrationUsecase_Register(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *usecase.RegistrationRequest
	}
	tests := []struct {
		name      string
		args      args
		wantToken string
		wantErr   bool
		mcFunc    func(registrator *mock.MockRegistrator, token *mock.MockTokenGenerator)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				params: &usecase.RegistrationRequest{
					Name:        "test",
					PhoneNumber: "08123112313",
				},
			},
			wantToken: "token",
			wantErr:   false,
			mcFunc: func(registrator *mock.MockRegistrator, token *mock.MockTokenGenerator) {
				registrator.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
				token.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("token", nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			registrator := mock.NewMockRegistrator(ctrl)
			tokenGenerator := mock.NewMockTokenGenerator(ctrl)
			r := usecase.NewRegistrationUsecase(registrator, tokenGenerator)

			tt.mcFunc(registrator, tokenGenerator)
			gotToken, err := r.Register(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegistrationUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("RegistrationUsecase.Register() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
