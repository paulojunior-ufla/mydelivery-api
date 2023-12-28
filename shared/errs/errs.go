package errs

type Type string

const (
	VALIDATION_ERROR  = "VALIDATION_ERROR"
	UNEXPECTED_ERROR  = "UNEXPECTED_ERROR"
	NOT_FOUND_ERROR   = "NOT_FOUND_ERROR"
	BAD_REQUEST_ERROR = "BAD_REQUEST_ERROR"
	CONFLICT_ERROR    = "CONFLICT_ERROR"
)

type Err struct {
	Message string
	ErrType Type
	Details []string
	Cause   error
}

func (e *Err) Error() string { return e.Message }

func NewValidationError(details []string) error {
	return &Err{
		Message: "um ou mais campos são inválidos",
		ErrType: VALIDATION_ERROR,
		Details: details,
	}
}

func NewUnexpectedError(err error) error {
	return &Err{
		Message: "ocorreu um erro inesperado",
		ErrType: UNEXPECTED_ERROR,
		Cause:   err,
	}
}

func NewBadRequestError(err error) error {
	return &Err{
		Message: "requisição inválida",
		ErrType: BAD_REQUEST_ERROR,
		Cause:   err,
	}
}

func NewNotFoundError(message string) error {
	return &Err{
		Message: message,
		ErrType: NOT_FOUND_ERROR,
	}
}

func NewConflictError(message string) error {
	return &Err{
		Message: message,
		ErrType: CONFLICT_ERROR,
	}
}
