package user

import (
	"context"
	"poc/internal/onboarding/entity"

	_ "github.com/lib/pq"
)

// 1 repository per aggregate root
type UserPostgresql struct {
	queries *Queries
}

func NewUserPostgresql(queries *Queries) *UserPostgresql {
	return &UserPostgresql{
		queries: queries,
	}
}

func (r *UserPostgresql) Register(ctx context.Context, user *entity.User) error {
	id, err := r.queries.InsertUser(ctx, &InsertUserParams{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
		CreatedAt:   user.CreatedAt,
	})
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func (r *UserPostgresql) GetUser(ctx context.Context, id int) (*entity.User, error) {
	userDB, err := r.queries.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:          id,
		Name:        userDB.Name,
		PhoneNumber: userDB.PhoneNumber,
		Role:        entity.UserRole(userDB.Role),
		CreatedAt:   userDB.CreatedAt,
	}, nil
}

func (r *UserPostgresql) GetUserAddresses(ctx context.Context, userID int) (*entity.User, error) {
	addresses, err := r.queries.GetUserAddresses(ctx, userID)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ID:        userID,
		Addresses: make([]*entity.Address, 0, len(addresses)),
	}
	for _, address := range addresses {
		user.Addresses = append(user.Addresses, &entity.Address{
			ID:        int(address.ID),
			City:      address.City,
			Street:    address.Street,
			ZipCode:   address.ZipCode,
			Latitude:  address.Latitude,
			Longitude: address.Longitude,
		})
	}

	return user, nil
}

func (r *UserPostgresql) AddAddress(ctx context.Context, userID int, address *entity.Address) error {
	id, err := r.queries.InsertAddress(ctx, &InsertAddressParams{
		UserID:    userID,
		Street:    address.Street,
		City:      address.City,
		ZipCode:   address.ZipCode,
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
	})

	if err != nil {
		return err
	}

	address.ID = id
	return nil
}
