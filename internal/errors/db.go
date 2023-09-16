package errors

// DatabaseError should be raised when db operation fails with unknown reason.
type DatabaseError struct{}

// NewDatabaseError creates a new instance of database error.
func NewDatabaseError() *DatabaseError {
	return &DatabaseError{}
}

// Error converts the error in string format.
func (e *DatabaseError) Error() string {
	return "The db operation failed with unknown reason. Please check logs to diagnose further."
}
