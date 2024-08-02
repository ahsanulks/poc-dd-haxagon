package entity

import (
	"regexp"
	"time"
)

const (
	phoneNumberMinimumLength = 10
	nameMinimumLength        = 3
)

func RegisterUser(name, phoneNumber string) (*User, error) {
	if err := nameValid(name); err != nil {
		return nil, err
	}

	if err := phoneNumberValid(phoneNumber); err != nil {
		return nil, err
	}

	return &User{
		Name:        name,
		PhoneNumber: phoneNumber,
		Role:        UserRoleNormal,
		CreatedAt:   time.Now(),
	}, nil
}

func nameValid(name string) error {
	if len(name) < nameMinimumLength {
		return ErrNameNotMatchMinimumLength
	}
	return nil
}

func phoneNumberValid(phoneNumber string) error {
	if len(phoneNumber) < phoneNumberMinimumLength {
		return ErrPhoneNumberNotMatchMinimumLength
	}

	if phoneNumber[0] != '0' {
		return ErrPhoneNumberNotStartWithZero
	}

	// check only containt number
	matched, err := regexp.MatchString(`^\d+$`, phoneNumber)
	if err != nil || !matched {
		return ErrPhoneNumberOnlyNumber
	}

	return nil
}
