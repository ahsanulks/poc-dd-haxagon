package entity

import businesserror "poc/internal/shared/business_error"

var (
	ErrProductRequired = businesserror.NewBusinessError(businesserror.InputValidationError, "product is required")
	ErrAddressRequired = businesserror.NewBusinessError(businesserror.InputValidationError, "address is required")
)
