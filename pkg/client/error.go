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
	_, ok := Cause(err).(ErrorResourceTypeNotFound)
	return ok
}

func Cause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
