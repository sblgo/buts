package sqldb

import (
	"database/sql"
	"errors"
)

type standard struct{}

func (s standard) Exec(db *sql.DB, stmt *Statement) (int, error) {
	switch stmt.Command {
	case CREATE_TABLE:
		return CreateTable(db, stmt)
	case INSERT:
		return InsertTable(db, stmt)
	}
	return 0, errors.New("undefined command: " + string(stmt.Command))
}

func (s standard) Query(db *sql.DB, stmt *Statement) (*sql.Rows, error) {
	return SelectTable(db, stmt)
}

var (
	standardDialect Dialect = standard{}
)

func init() {
	Register("standard", standardDialect)
}
