package apperrors

import "errors"

var (
	ErrNumOfDistances      = errors.New("number-of-distances-mistmatch")
	ErrMessageLenMismatch  = errors.New("message-len-mistmatch")
	ErrDecodingMessage     = errors.New("error-decoding-message")
	ErrUnknowMessageSource = errors.New("unknow-satellite")
)
