package postgres

import (
	"errors"

	_ "github.com/lib/pq"
)

var ErrQueryPreparation = errors.New("query preparation error")
