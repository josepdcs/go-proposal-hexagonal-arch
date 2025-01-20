package errors

import "github.com/pkg/errors"

// User errors

// ErrUserNotFound is an error returned when a user is not found.
var ErrUserNotFound = errors.New("user not found")

// ErrUserAlreadyExists is an error returned when a user already exists.
var ErrUserAlreadyExists = errors.New("user already exists")
