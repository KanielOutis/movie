package repository

import "errors"

// ErrorNotFound is returned when a metadata is not found.
var ErrorNotFound = errors.New("not found")
