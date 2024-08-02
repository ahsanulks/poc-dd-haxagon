package repository

import (
	"context"
	"poc/internal/order/entity"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderPostgresql struct {
	conn         *pgxpool.Pool
	queryBuilder sq.StatementBuilderType
}

func NewOrderPostgresql(conn *pgxpool.Pool) *OrderPostgresql {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &OrderPostgresql{conn, psql}
}

func (r *OrderPostgresql) CreateOrder(ctx context.Context, order *entity.Order) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	sql, args, err := r.queryBuilder.Insert("orders").Columns("user_id", "total_price", "created_at").
		Values(order.UserID, order.TotalPrice, order.CreatedAt).Suffix("RETURNING id").ToSql()
	if err != nil {
		return err
	}

	var orderID int64
	err = tx.QueryRow(ctx, sql, args...).Scan(&orderID)
	if err != nil {
		tx.Rollback(ctx)
	}
	sql, args, err = r.queryBuilder.Insert("order_addresses").Columns("order_id", "street", "city", "zip_code", "latitude", "longitude").
		Values(orderID, order.Address.Street, order.Address.City, order.Address.ZipCode, order.Address.Latitude, order.Address.Longitude).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
	}
	order.ID = int(orderID)
	return tx.Commit(ctx)
}
