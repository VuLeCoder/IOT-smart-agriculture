package services

import "errors"

var (
	ErrInvalidNumber      = errors.New("number must be between 1 and 100")
	ErrInvalidCredentials = errors.New("email or password is incorrect")
)
