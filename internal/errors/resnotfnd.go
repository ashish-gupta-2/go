package errors

import (
	"fmt"
)

// ResourceNotFoundError should be raised when the specified resource not available.
type ResourceNotFoundError struct {
	Type string
	Name string
}

// NewResourceNotFoundError creates a new instance of 'resource not found' error.
func NewResourceNotFoundError(t, name string) *ResourceNotFoundError {
	return &ResourceNotFoundError{
		Type: t,
		Name: name,
	}
}

// Error converts the error in string format.
func (e *ResourceNotFoundError) Error() string {
	f := "The %s resource '%s' was not found. Please use different resource name."
	return fmt.Sprintf(f, e.Type, e.Name)
}
