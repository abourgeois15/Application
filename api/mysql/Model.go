package mysql

import "database/sql"

type Model struct {
	Db *sql.DB
}
