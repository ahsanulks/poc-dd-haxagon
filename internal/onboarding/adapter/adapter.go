package adapter

import (
	"context"
	"poc/internal/onboarding/entity"
	"poc/internal/onboarding/usecase"
)

type RegistrationUsecase interface {
	Register(ctx context.Context, params *usecase.RegistrationRequest) (token string, err error)
}

type AddressManagementUsecase interface {
	AddAddress(ctx context.Context, addressReq *usecase.AddAddressRequest) (*entity.Address, error)
	GetAddress(ctx context.Context, userID int) ([]*entity.Address, error)
}
