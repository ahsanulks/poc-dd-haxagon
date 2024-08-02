package entity

import "errors"

var (
	ErrPhoneNumberOnlyNumber            = errors.New("phone number should only contains number")
	ErrPhoneNumberNotMatchMinimumLength = errors.New("phone number should have minimum length of 10")
	ErrPhoneNumberNotStartWithZero      = errors.New("phone number should start with 0")
	ErrNameNotMatchMinimumLength        = errors.New("name should have minimum length of 3")
	ErrAddressStreetRequired            = errors.New("address street is required")
	ErrAddressCityRequired              = errors.New("address city is required")
	ErrAddressZipCodeRequired           = errors.New("address zip code is required")
	ErrUserAddressExceededLimit         = errors.New("user address exceeded limit")
)
