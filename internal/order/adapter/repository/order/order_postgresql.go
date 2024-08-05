package order

import (
	"context"
	"poc/internal/order/entity"

	"github.com/jackc/pgx/v5"
)

type DBTransaction interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type OrderPostgresql struct {
	conn    DBTransaction
	queries *Queries
}

func NewOrderPostgresql(conn DBTransaction, queries *Queries) *OrderPostgresql {
	return &OrderPostgresql{conn, queries}
}

func (r *OrderPostgresql) CreateOrder(ctx context.Context, order *entity.Order) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	orderID, err := r.queries.WithTx(tx).InsertOrder(ctx, &InsertOrderParams{
		UserID:     order.UserID,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
	})
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	_, err = r.queries.WithTx(tx).InsertOrderAddress(ctx, &InsertOrderAddressParams{
		OrderID:   orderID,
		Street:    order.Address.Street,
		City:      order.Address.City,
		ZipCode:   order.Address.ZipCode,
		Latitude:  order.Address.Latitude,
		Longitude: order.Address.Longitude,
	})
	order.ID = orderID
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}
