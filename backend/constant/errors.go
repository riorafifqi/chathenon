package constant

import "errors"

var (
	ErrDBConnection     = errors.New("database connection error")
	ErrDBMigration      = errors.New("database migration error")
	ErrDBQuery          = errors.New("database query error")
	ErrRecordNotFound   = errors.New("record not found")
	ErrEmailAlreadyUsed = errors.New("email already in use")
)