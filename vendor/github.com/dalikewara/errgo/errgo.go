package errgo

import "errors"

type ErrGo interface {
	Error() error
	GetError() error
	GetCode() string
	GetMessage() string
	GetStatus() int
}

type errgo struct {
	code    string
	message string
	status  int
	error   error
}

// New generates new errgo that implements ErrGo.
func New(code, message string) ErrGo {
	return &errgo{
		code:    code,
		message: message,
		error:   generateError(code, message, "||"),
	}
}

// NewWithStatus generates new errgo with status that implements ErrGo.
func NewWithStatus(code, message string, status int) ErrGo {
	return &errgo{
		code:    code,
		message: message,
		status:  status,
		error:   generateError(code, message, "||"),
	}
}

// Error gets errgo error.
func (e *errgo) Error() error {
	return e.error
}

// GetError gets errgo error.
func (e *errgo) GetError() error {
	return e.error
}

// GetCode gets errgo code.
func (e *errgo) GetCode() string {
	return e.code
}

// GetMessage gets errgo message.
func (e *errgo) GetMessage() string {
	return e.message
}

// GetStatus gets errgo status.
func (e *errgo) GetStatus() int {
	return e.status
}

// generateError generates new error.
func generateError(code, message, separator string) error {
	return errors.New(code + separator + message)
}
