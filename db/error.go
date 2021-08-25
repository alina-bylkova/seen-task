package db

type DbError struct {
	originalError error
}

// Error returns custom DbError
func (e *DbError) Error() string {
	return e.originalError.Error()
}
