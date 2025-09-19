package customerr

type DatabaseError struct {
	originalErr error
	message     string
}

func NewDatabaseError(err error) error {
	return &DatabaseError{
		originalErr: err,
		message:     "database error",
	}
}

func (e *DatabaseError) Error() string {
	return e.message
}

func (e *DatabaseError) Unwrap() error {
	return e.originalErr
}
