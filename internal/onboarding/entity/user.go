package entity

import "time"

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleNormal UserRole = "user"
)

type User struct {
	ID          int
	Name        string
	PhoneNumber string
	Role        UserRole
	CreatedAt   time.Time
	Addresses   []*Address
}
