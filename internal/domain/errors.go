package domain

import "errors"

var (
	ErrNotFound           = errors.New("resource not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrConflict           = errors.New("resource conflict")
	ErrInternal           = errors.New("internal server error")
	ErrDuplicateUsername  = errors.New("username already exists")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrAlreadyFollowing   = errors.New("already following")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
