package user

import (
	"context"
	"errors"
	"poc/internal/onboarding/adapter/repository/user"
	"poc/internal/onboarding/entity"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pashagolub/pgxmock/v4"
)

func TestUserPostgresql_GetUserAddresses(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
		mcFunc  func(mock pgxmock.PgxPoolIface)
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &entity.User{
				ID: 1,
				Addresses: []*entity.Address{
					{
						ID:        1,
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
				mock.ExpectQuery("-- name: GetUserAddresses :many").
					WithArgs(1).
					WillReturnRows(mock.NewRows([]string{"id", "user_id", "street", "city", "zip", "latitude", "longitude"}).
						AddRow(1, 1, "street", "city", "zip", 1.1, 1.1))
			},
		},
		{
			name: "error db conn",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want:    nil,
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectQuery("-- name: GetUserAddresses :many").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
		},
		{
			name: "error scan",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want:    nil,
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectQuery("-- name: GetUserAddresses :many").
					WithArgs(1).
					WillReturnRows(mock.NewRows([]string{"id", "user_id", "street", "city", "zip", "latitude", "longitude"}).
						AddRow(1, 1, "street", "city", "zip", "1.1", "1.1"))
			},
		},
		{
			name: "error rows",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want:    nil,
			wantErr: true,
			mcFunc: func(mock pgxmock.PgxPoolIface) {
				mock.ExpectQuery("-- name: GetUserAddresses :many").
					WithArgs(1).
					WillReturnRows(mock.NewRows([]string{"id", "user_id", "street", "city", "zip", "latitude", "longitude"}).
						RowError(0, errors.New("row error")))
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

			r := user.NewUserPostgresql(user.New(mock))
			got, err := r.GetUserAddresses(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPostgresql.GetUserAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPostgresql.GetUserAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
