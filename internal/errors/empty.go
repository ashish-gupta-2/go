package errors

import "fmt"

// EmptyError should be raised when the value is empty.
type EmptyError struct {
	Key string
}

// NewEmptyError creates a new instance of empty error.
func NewEmptyError(key string) *EmptyError {
	return &EmptyError{
		Key: key,
	}
}

// Error converts the error in string format.
func (e *EmptyError) Error() string {
	return fmt.Sprintf("The %s is empty. Please input valid data.", e.Key)
}
