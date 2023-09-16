package errors

import "fmt"

// UnMarshalError should be raised when the value is not valid.
type UnMarshalError struct {
	Key string
}

// NewUnMarshalError creates a new instance of unmarshal error.
func NewUnMarshalError(key string) *UnMarshalError {
	return &UnMarshalError{
		Key: key,
	}
}

// Error converts the error in string format.
func (e *UnMarshalError) Error() string {
	return fmt.Sprintf("The %s can't be unmarshalled. Please input valid data.", e.Key)
}
