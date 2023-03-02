package errors

import (
	"errors"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("record not found")
	ErrDataValidation  = errors.New("data validation failed")
)
