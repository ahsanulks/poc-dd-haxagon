package entity

import "errors"

var (
	ErrProductRequired = errors.New("product is required")
	ErrAddressRequired = errors.New("address is required")
)
