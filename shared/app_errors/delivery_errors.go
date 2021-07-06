package apperrors

import "errors"

var (
	ErrBadInput = errors.New("wrong-req-body")
)

type DeliveryError struct {
	Message string
}

func NewDeliveryError(err error) DeliveryError {
	return DeliveryError{
		Message: err.Error(),
	}
}
