package errs

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func New(code int, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
