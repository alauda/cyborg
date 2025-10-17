package client

import "errors"

type ErrorResourceTypeNotFound struct {
	message string
}

func NewTypeNotFoundError(message string) ErrorResourceTypeNotFound {
	return ErrorResourceTypeNotFound{message: message}
}

func (e ErrorResourceTypeNotFound) Error() string {
	return e.message
}

func IsResourceTypeNotFound(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(ErrorResourceTypeNotFound); ok {
		return true
	} else {
		unwrapErr := errors.Unwrap(err)
		if unwrapErr == nil {
			return false
		} else if _, ok := unwrapErr.(ErrorResourceTypeNotFound); ok {
			return true
		}
	}
	return false
}
