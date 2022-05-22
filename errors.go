package vcard

import "errors"

var (
	// ErrUnsupportedType indicates that an unsupported interface type was sent for conversion.
	ErrUnsupportedType = errors.New("unsupported type")
)
