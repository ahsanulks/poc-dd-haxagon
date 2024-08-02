package usecase

import (
	"context"
	"poc/internal/onboarding/entity"
)

type AddressUsecase struct {
	modifier entity.AddressModifier
	getter   entity.AddressGetter
}

func NewAddressUsecase(
	modifier entity.AddressModifier,
	getter entity.AddressGetter,
) *AddressUsecase {
	return &AddressUsecase{
		modifier: modifier,
		getter:   getter,
	}
}

func (a *AddressUsecase) GetAddress(ctx context.Context, userID int) ([]*entity.Address, error) {
	user, err := a.getter.GetUserAddresses(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user.Addresses, nil
}

func (a *AddressUsecase) AddAddress(ctx context.Context, addressReq *AddAddressRequest) (*entity.Address, error) {
	addressBuilder := new(entity.AddressBuilder)
	address, err := addressBuilder.City(addressReq.City).
		Street(addressReq.Street).
		ZipCode(addressReq.ZipCode).
		Latitude(addressReq.Latitude).
		Longitude(addressReq.Longitude).
		Build()

	if err != nil {
		return nil, err
	}

	user, err := a.getter.GetUserAddresses(ctx, addressReq.UserID)
	if err != nil {
		return nil, err
	}

	if err := user.AddAddress(address); err != nil {
		return nil, err
	}

	err = a.modifier.AddAddress(ctx, user.ID, address)
	if err != nil {
		return nil, err
	}
	return address, nil
}
