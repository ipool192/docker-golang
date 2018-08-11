package interfaces

import "database/sql"

// ISQLConnection - initialize interface for SQL Connection
type ISQLConnection interface {
	Connect() (*sql.DB, error)
}
