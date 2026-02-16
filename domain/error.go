package domain

import "net/http"

type Error struct {
	StatusCode int
	Err        error
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewInternalServerError(err error) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

func NewBadRequestError(err error) *Error {
	return &Error{
		StatusCode: http.StatusBadRequest,
		Err:        err,
	}
}

func NewUnauthorizedError(err error) *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		Err:        err,
	}
}

func NewForbiddenError(err error) *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Err:        err,
	}
}
