package internal

import "database/sql"

type AppContext struct {
	DB *sql.DB
}
