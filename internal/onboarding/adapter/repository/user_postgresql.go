package repository

import (
	"context"
	"poc/internal/onboarding/entity"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

// 1 repository per aggregate root
type UserPostgresql struct {
	conn         *pgxpool.Pool
	queryBuilder sq.StatementBuilderType
}

func NewUserPostgresql(conn *pgxpool.Pool) *UserPostgresql {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &UserPostgresql{
		conn:         conn,
		queryBuilder: psql,
	}
}

func (r *UserPostgresql) Register(ctx context.Context, user *entity.User) error {
	sql, args, err := r.queryBuilder.Insert("users").Columns("name", "phone_number", "role", "created_at").
		Values(user.Name, user.PhoneNumber, user.Role, user.CreatedAt).Suffix("RETURNING id").ToSql()
	if err != nil {
		return err
	}

	var id int64
	err = r.conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *UserPostgresql) GetUser(ctx context.Context, id int) (*entity.User, error) {
	sql, args, err := r.queryBuilder.Select("id", "name", "phone_number", "role", "created_at").
		From("users").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	row := r.conn.QueryRow(ctx, sql, args...)
	user := new(entity.User)
	err = row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserPostgresql) GetUserAddresses(ctx context.Context, userID int) (*entity.User, error) {
	sql, args, err := r.queryBuilder.Select("a.id", "a.city", "a.street", "a.zip_code", "a.latitude", "a.longitude").
		From("users u").Join("addresses a on u.id = a.user_id").Where(sq.Eq{"u.id": userID}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &entity.User{
		ID: userID,
	}
	user.Addresses = []*entity.Address{}
	for rows.Next() {
		address := new(entity.Address)
		err = rows.Scan(&address.ID, &address.City, &address.Street, &address.ZipCode, &address.Latitude, &address.Longitude)
		if err != nil {
			return nil, err
		}
		user.Addresses = append(user.Addresses, address)
	}

	return user, nil
}

func (r *UserPostgresql) AddAddress(ctx context.Context, userID int, address *entity.Address) error {
	sql, args, err := r.queryBuilder.Insert("addresses").Columns("city", "street", "zip_code", "latitude", "longitude", "user_id").
		Values(address.City, address.Street, address.ZipCode, address.Latitude, address.Longitude, userID).Suffix("RETURNING id").ToSql()
	if err != nil {
		return err
	}

	var id int64
	err = r.conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return err
	}

	address.ID = int(id)
	return nil
}
