package errors

import "errors"

func NewInternalError() error {
	return &BizError{
		error: errors.New("internal Error"),
		code:  InternalError,
	}
}

func NewError(msg string, code int32) error {
	return &BizError{
		error: errors.New(msg),
		code:  code,
	}
}

type BizError struct {
	error
	code int32
}

func (i *BizError) GetCode() int32 {
	return i.code
}
