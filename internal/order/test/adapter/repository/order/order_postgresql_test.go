package order_test

import (
	"context"
	"errors"
	"poc/internal/order/adapter/repository/order"
	"poc/internal/order/entity"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v4"
)

func TestOrderPostgresql_CreateOrder(t *testing.T) {
	now := time.Now()
	type args struct {
		ctx   context.Context
		order *entity.Order
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mcFunc  func(mock pgxmock.PgxPoolIface)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				order: &entity.Order{
					UserID:     1,
					TotalPrice: 100,
					CreatedAt:  now,
					Address: &entity.Address{
						Street:    "street",
						City:      "city",
						ZipCode:   "zip",
						Latitude:  1.1,
						Longitude: 1.1,
					},
				},
			},
			wantErr: false,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders").
					WithArgs(1, float64(100), now).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery("INSERT INTO order_address").
					WithArgs(1, "street", "city", "zip", 1.1, 1.1).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error begin",
			args: args{
				ctx: context.Background(),
				order: &entity.Order{
					UserID:     1,
					TotalPrice: 100,
				},
			},
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin().WillReturnError(errors.New("begin error"))
			},
		},
		{
			name: "error insert order",
			args: args{
				ctx: context.Background(),
				order: &entity.Order{
					UserID:     1,
					TotalPrice: 100,
					CreatedAt:  now,
				},
			},
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders").
					WithArgs(1, float64(100), now).
					WillReturnError(errors.New("insert order error"))
				mock.ExpectRollback()
			},
		},
		{
			name: "error insert order address",
			args: args{
				ctx: context.Background(),
				order: &entity.Order{
					UserID:     1,
					TotalPrice: 100,
					CreatedAt:  now,
					Address: &entity.Address{
						Street:    "street",
						City:      "city",
						ZipCode:   "zip",
						Latitude:  1.1,
						Longitude: 1.1,
					},
				},
			},
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders").
					WithArgs(1, float64(100), now).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery("INSERT INTO order_address").
					WithArgs(1, "street", "city", "zip", 1.1, 1.1).
					WillReturnError(errors.New("insert order address error"))
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()
			tt.mcFunc(mock)
			r := order.NewOrderPostgresql(mock, order.New(mock))
			if err := r.CreateOrder(tt.args.ctx, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("OrderPostgresql.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
