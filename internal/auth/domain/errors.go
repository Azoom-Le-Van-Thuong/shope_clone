package domain

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrTokenExpired      = errors.New("token has expired")
	ErrTokenInvalid      = errors.New("invalid token")
)
