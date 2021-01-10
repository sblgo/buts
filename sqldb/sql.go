package sqldb

import (
	"database/sql"
	"errors"
	"log"

	"github.com/sblgo/buts"
)

//
type Command int

const (
	INSERT Command = iota + 1
	UPDATE
	DELETE
	SELECT
	CREATE_TABLE
)

type Operator string

const (
	OP_EQ Operator = " = "
)

type Column struct {
	Name     string
	Value    interface{}
	Operator Operator
}

type BaseType string

const (
	BT_VARCHAR BaseType = "VARCHAR"
	BT_INT              = "INT"
)

type ColumnDef struct {
	Type       BaseType
	Length     int
	Fraction   int
	Default    string
	PrimaryKey bool
}

type Columns []Column

type Statement struct {
	Table        string
	Command      Command
	Presentation Columns
	Condition    Columns
}

type Dialect interface {
	Exec(db *sql.DB, stmt *Statement) (int, error)
	Query(db *sql.DB, stmt *Statement) (*sql.Rows, error)
}

var (
	dialects map[string]Dialect = make(map[string]Dialect)
)

func Register(name string, dialect Dialect) {
	if _, ok := dialects[name]; !ok {
		dialects[name] = dialect
	}
}

type typeSystem struct {
	connection  *sql.DB
	dialectName string
	dialect     Dialect
}

var (
	dbprep map[*sql.DB]bool = make(map[*sql.DB]bool)
)

func New(dialect string, db *sql.DB) (buts.TypeSystem, error) {
	ts := &typeSystem{
		dialectName: dialect,
		connection:  db,
	}
	if ts.dialect = dialects[dialect]; ts.dialect == nil {
		return nil, errors.New("unknown dialect: " + dialect)
	}
	if prep, ok := dbprep[db]; !ok || !prep {
		if err := ts.prepareTypeSystemTables(); err != nil {
			return nil, err
		} else {
			dbprep[db] = true
		}

	}
	return ts, nil
}

func (ts *typeSystem) New(kind buts.Kind, name string) buts.Type {
	switch kind {
	case buts.Element:
		return ts.newElement(name)
	}
	return nil
}

func (ts *typeSystem) prepareTypeSystemTables() error {
	for _, stmt := range tableDefinitions {
		existStat := Statement{
			Table:   stmt.Table,
			Command: SELECT,
			Presentation: Columns{
				{
					Name: "count(*)",
				},
			},
		}
		if _, err := ts.dialect.Query(ts.connection, &existStat); err != nil {
			_, err := ts.dialect.Exec(ts.connection, &stmt)
			if err != nil {
				return err
			}
		}
	}
	for _, s := range insDatElements {
		stmt := Statement{
			Table:   tabDatElement.Table,
			Command: INSERT,
			Presentation: Columns{
				{Name: "NAME", Value: &s.Name},
				{Name: "COL_LENGTH", Value: &s.ColLength},
				{Name: "COL_FRAC", Value: &s.ColFrac},
				{Name: "DB_TYPE", Value: &s.DbType},
				{Name: "GO_TYPE", Value: &s.GoType},
				{Name: "DESCRIPTION", Value: &s.Description},
			},
		}
		i, err := ts.dialect.Exec(ts.connection, &stmt)
		log.Printf(" %d - %v \n", i, err)
	}
	return nil
}

func (ts *typeSystem) newElement(name string) buts.Type {
	stmt := Statement{
		Table:        tabDatElement.Table,
		Command:      SELECT,
		Presentation: tabDatElement.Presentation,
		Condition: Columns{
			{
				Name:     tabDatElement.Presentation[0].Name,
				Value:    &name,
				Operator: OP_EQ,
			},
		},
	}
	if rows, err := ts.dialect.Query(ts.connection, &stmt); err != nil {
		return nil
	} else if rows.Next() {
		te := new(typeElement)
		err = rows.Scan(&te.name, &te.colLength, &te.colFrac, &te.dbType, &te.goType, &te.description)
		if err != nil {
			return nil
		} else {
			return te
		}
	}
	return nil
}

type typeElement struct {
	name, description                  string
	colLength, colFrac, dbType, goType int
}

func (te *typeElement) New() buts.Value {
	return buts.Value{}
}
