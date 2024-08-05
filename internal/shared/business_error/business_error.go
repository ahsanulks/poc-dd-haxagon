package businesserror

type ErrorCode string

const (
	InputValidationError ErrorCode = "E001"
	DataValidationError  ErrorCode = "E002"
	NotFound             ErrorCode = "E404"
	InternalError        ErrorCode = "E500"
)

type BusinessError struct {
	code    ErrorCode
	message string
}

func NewBusinessError(code ErrorCode, message string) *BusinessError {
	return &BusinessError{
		code:    code,
		message: message,
	}
}

func (e BusinessError) Error() string {
	return e.message
}

func (e BusinessError) Code() ErrorCode {
	return e.code
}
