package errors

import "errors"

var (
	ErrTokenWasNotProvided  = errors.New("auth token was not provided")
	ErrInvalidSigningMethod = errors.New("provided token was signed via invalid method")
	ErrInternal             = errors.New("internal server error")
	ErrInvalidToken         = errors.New("provided token is invalid")
	ErrTokenExpired         = errors.New("token has expired")
)