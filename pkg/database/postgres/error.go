package postgres

import "fmt"

// PersistenceError is a generic database error.
type PersistenceError struct {
	Message string `json:"message"`
}

// NewPreparationError represents when a query return an error on prepation.
func NewPreparationError(queryName string, repository string, err error) *PersistenceError {
	preparationErr := fmt.Errorf(
		"unable to prepare the query:`%s` on %s repository, original err: %s",
		queryName,
		repository,
		err.Error(),
	)
	return newPersistenceError(preparationErr, "prepare", "postgres")
}

// NewStatementNotPreparedError represents when a query return an error trying to access a not prepared query.
func NewStatementNotPreparedError(queryName string, repository string) *PersistenceError {
	preparationErr := fmt.Errorf("query `%s` is not prepared on %s repository", queryName, repository)
	return newPersistenceError(preparationErr, "statement not prepared", "postgres")
}

func (e *PersistenceError) Error() string {
	return e.Message
}

func newPersistenceError(originalErr error, action, datasource string) *PersistenceError {
	return &PersistenceError{
		Message: fmt.Sprintf("%s persistence error on `%s`: %s", datasource, action, originalErr.Error()),
	}
}
