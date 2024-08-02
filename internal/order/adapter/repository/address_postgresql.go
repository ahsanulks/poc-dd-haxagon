package repository

import (
	"context"
	"poc/internal/order/entity"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressPostgresql struct {
	conn         *pgxpool.Pool
	queryBuilder sq.StatementBuilderType
}

func NewAddressPostgresql(conn *pgxpool.Pool) *AddressPostgresql {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &AddressPostgresql{conn, psql}
}

func (r *AddressPostgresql) GetUserAddresses(ctx context.Context, userID, addressID int) (*entity.Address, error) {
	sql, args, err := r.queryBuilder.Select("id", "street", "city", "zip_code", "latitude", "longitude").
		From("addresses").Where(sq.Eq{"user_id": userID, "id": addressID}).ToSql()
	if err != nil {
		return nil, err
	}

	var address entity.Address
	err = r.conn.QueryRow(ctx, sql, args...).Scan(&address.ID, &address.Street, &address.City, &address.ZipCode, &address.Latitude, &address.Longitude)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
