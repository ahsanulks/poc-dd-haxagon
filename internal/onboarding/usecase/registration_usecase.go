package usecase

import (
	"context"
	"poc/internal/onboarding/entity"
)

type RegistrationUsecase struct {
	registrator    entity.Registrator
	tokenGenerator entity.TokenGenerator
}

func NewRegistrationUsecase(
	registrator entity.Registrator,
	tokenGenerator entity.TokenGenerator,
) *RegistrationUsecase {
	return &RegistrationUsecase{
		registrator:    registrator,
		tokenGenerator: tokenGenerator,
	}
}

// please use the business term for the method name. instead of create user better is register or registration
func (r *RegistrationUsecase) Register(ctx context.Context, params *RegistrationRequest) (token string, err error) {
	user, err := entity.RegisterUser(params.Name, params.PhoneNumber)
	if err != nil {
		return token, err
	}

	if err := r.registrator.Register(ctx, user); err != nil {
		return token, err
	}

	return r.tokenGenerator.GenerateToken(ctx, user)
}
