package http

import (
	"errors"
)

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrMissingAuthToken = errors.New("missing auth token")
	ErrInvalidAuthToken = errors.New("invalid auth token")
	ErrWrongAuthToken   = errors.New("wrong auth token")
	ErrValidAuthToken   = errors.New("token is not valid")
	ErrInvalidUserID    = errors.New("invalid user id")
)
