package driver

import "database/sql"

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}
