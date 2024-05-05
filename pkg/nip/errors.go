package nip

import "errors"

var (
	// ErrEmptyRule is returned when a rule is empty
	ErrEmptyRule = errors.New("empty rule")
)
