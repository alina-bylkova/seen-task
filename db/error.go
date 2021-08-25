package db

// DbError is a custom error that shouldn't be sent out to clients
type DbError struct {
	originalError error
}

// Error returns custom DbError
func (e *DbError) Error() string {
	return e.originalError.Error()
}
