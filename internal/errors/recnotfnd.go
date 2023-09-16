package errors

import (
	"fmt"
)

// RecordNotFoundError should be raised when the specified db record not available.
type RecordNotFoundError struct {
	Entity   string
	Criteria string
}

// NewRecordNotFoundError creates a new instance of 'record not found' error.
func NewRecordNotFoundError(entity, criteria string) *RecordNotFoundError {
	return &RecordNotFoundError{
		Entity:   entity,
		Criteria: criteria,
	}
}

// Error converts the error in string format.
func (e *RecordNotFoundError) Error() string {
	f := "The %s record searched by '%s' was not found."
	return fmt.Sprintf(f, e.Entity, e.Criteria)
}
