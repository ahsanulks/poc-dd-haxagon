package entity

import (
	businesserror "poc/internal/shared/business_error"
)

var (
	ErrPhoneNumberOnlyNumber            = businesserror.NewBusinessError(businesserror.InputValidationError, "phone number should only contains number")
	ErrPhoneNumberNotMatchMinimumLength = businesserror.NewBusinessError(businesserror.InputValidationError, "phone number should have minimum length of 10")
	ErrPhoneNumberNotStartWithZero      = businesserror.NewBusinessError(businesserror.InputValidationError, "phone number should start with 0")
	ErrNameNotMatchMinimumLength        = businesserror.NewBusinessError(businesserror.InputValidationError, "name should have minimum length of 3")
	ErrAddressStreetRequired            = businesserror.NewBusinessError(businesserror.InputValidationError, "address street is required")
	ErrAddressCityRequired              = businesserror.NewBusinessError(businesserror.InputValidationError, "address city is required")
	ErrAddressZipCodeRequired           = businesserror.NewBusinessError(businesserror.InputValidationError, "address zip code is required")
	ErrUserAddressExceededLimit         = businesserror.NewBusinessError(businesserror.DataValidationError, "user address exceeded limit")
)
