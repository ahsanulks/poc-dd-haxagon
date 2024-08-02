package entity

import "context"

type Registrator interface {
	Register(ctx context.Context, user *User) error
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, user *User) (string, error)
}

type AddressModifier interface {
	AddAddress(ctx context.Context, userID int, address *Address) error
}

type AddressGetter interface {
	GetUserAddresses(ctx context.Context, userID int) (*User, error)
}
