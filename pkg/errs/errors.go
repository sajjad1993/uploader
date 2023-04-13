package errs

import (
	"errors"
	"fmt"
)

var (
	ErrValidation     = errors.New("2|validation error")
	ErrNotFound       = errors.New("3|not found error")
	ErrUploadFailed   = errors.New("4|upload failed")
	ErrConflictEntity = errors.New("6|conflict entity")
	ErrUnprocessable  = errors.New("8|unprocessable entity")
	ErrCacheMis       = errors.New("9|key not found")
)

func NewValidationError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrValidation, errMsg)
}

func NewNotFoundError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrNotFound, errMsg)
}

func NewUploadFailed(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrUploadFailed, errMsg)
}

func NewConflictEntity(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrConflictEntity, errMsg)
}

func NewUnprocessable(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrUnprocessable, errMsg)
}

func NewErrCacheMis(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrCacheMis, errMsg)
}
