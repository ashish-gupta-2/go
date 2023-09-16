package errors

import "fmt"

// FormatError should be raised when the input format is invalid.
type FormatError struct {
	Key   string
	Value string
}

// NewFormatError creates a new instance of format error.
func NewFormatError(key, value string) *FormatError {
	return &FormatError{
		Key:   key,
		Value: value,
	}
}

// Error converts the error in string format.
func (e *FormatError) Error() string {
	f := "The %s '%s' is in invalid format. Please supply data in valid format."
	return fmt.Sprintf(f, e.Key, e.Value)
}
