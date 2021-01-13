package sqldb

import (
	"database/sql"
	"errors"
	"log"
	"reflect"

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
	Sort         Columns
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
	case buts.Structure:
		return ts.newStructure(name)
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
	return ts.Register(insDatElements, insDatStructure, nil)
}

func (ts *typeSystem) newElement(name string) buts.Type {
	if be, ok := ts.readElement(name); ok {
		te := &typeElement{
			typeNil: typeNil{
				typeSystem:  ts,
				kind:        buts.Element,
				name:        be.Name,
				description: be.Description,
				reflGoType:  goTypeMap[be.GoType],
			},
			goType:         be.GoType,
			dbType:         be.DbType,
			dbLength:       be.DbLength,
			dbDecimals:     be.DbDecimals,
			tags:           be.Tags,
			domain:         be.Domain,
			domainTable:    be.DomainTable,
			domainGoColumn: be.DomainGoColumn,
			domainDbColumn: be.DomainDbColumn,
			conversion:     be.Conversion,
			reflDbType:     dbTypeMap[be.DbType],
		}
		return te
	} else {
		return nil
	}

}

func (ts *typeSystem) newStructure(name string) buts.Type {
	if se, ok := ts.readStructure(name); ok {
		strct := &typeStructure{
			typeNil: typeNil{
				typeSystem: ts,
				kind:       buts.Structure,
				name:       se.Name,
			},
			Description: se.Description,
			Tags:        se.Tags,
			fields:      make([]typeField, 0),
			reflFields:  make([]reflect.StructField, 0),
		}

		for _, sef := range se.Items {
			field := typeField{
				Name:        sef.Name,
				Description: sef.Description,
				Kind:        sef.Kind,
				Type:        sef.Type,
			}
			field.fieldType = ts.New(field.Kind, field.Type)
			sf := reflect.StructField{
				Name: field.Name,
				Type: field.fieldType.ReflGoType(),
			}
			strct.fields = append(strct.fields, field)
			strct.reflFields = append(strct.reflFields, sf)

		}
		strct.reflGoType = reflect.StructOf(strct.reflFields)
		return strct
	} else {
		return nil
	}
}

type RegisterError struct {
}

func (ts *typeSystem) Register(elms []buts.ElementReg, strts []buts.StructureReg, tabs []buts.TableReg) error {
	for _, s := range elms {
		stmt := Statement{
			Table:   tabDatElement.Table,
			Command: INSERT,
			Presentation: Columns{
				{Name: "NAME", Value: &s.Name},
				{Name: "DESCRIPTION", Value: &s.Description},
				{Name: "GO_TYPE", Value: &s.GoType},
				{Name: "DB_TYPE", Value: &s.DbType},
				{Name: "COL_LENGTH", Value: &s.DbLength},
				{Name: "COL_FRAC", Value: &s.DbDecimals},
				{Name: "TAGS", Value: &s.Tags},
				{Name: "DOMAIN_NAME", Value: &s.Domain},
				{Name: "DOMAIN_TABLE", Value: &s.DomainTable},
				{Name: "DOMAIN_GOCOL", Value: &s.DomainGoColumn},
				{Name: "DOMAIN_DBCOL", Value: &s.DomainDbColumn},
				{Name: "DOMAIN_CONV", Value: &s.Conversion},
			},
		}
		i, err := ts.dialect.Exec(ts.connection, &stmt)
		log.Printf(" %d - %v \n", i, err)
	}
	for _, s := range strts {
		stmt := Statement{
			Table:   tabDatStructure.Table,
			Command: INSERT,
			Presentation: Columns{
				{Name: "NAME", Value: &s.Name},
				{Name: "DESCRIPTION", Value: &s.Description},
				{Name: "TAGS", Value: &s.Tags},
			},
		}
		i, err := ts.dialect.Exec(ts.connection, &stmt)
		if err != nil {
			log.Printf(" %d - %v \n", i, err)
		}
		for idx, f := range s.Items {
			stmt := Statement{
				Table:   tabDatFeld.Table,
				Command: INSERT,
				Presentation: Columns{
					{Name: "STRUCT_NAME", Value: &s.Name},
					{Name: "POS", Value: &idx},
					{Name: "NAME", Value: &f.Name},
					{Name: "DESCRIPTION", Value: &f.Description},
					{Name: "KIND", Value: &f.Kind},
					{Name: "TYPE", Value: &f.Type},
				},
			}
			i, err := ts.dialect.Exec(ts.connection, &stmt)
			if err != nil {
				log.Printf(" %d - %v \n", i, err)
			}
		}
	}

	for _, t := range tabs {
		stmt := Statement{
			Table:   tabDatTable.Table,
			Command: INSERT,
			Presentation: Columns{
				{Name: "NAME", Value: &t.Name},
				{Name: "DESCRIPTION", Value: &t.Description},
				{Name: "TAGS", Value: &t.Tags},
				{Name: "TABLE_TYPE", Value: &t.TableType},
				{Name: "KIND", Value: &t.Kind},
				{Name: "TYPE", Value: &t.Type},
				{Name: "SQL_NAME", Value: &t.SQLName},
			},
		}
		i, err := ts.dialect.Exec(ts.connection, &stmt)
		if err != nil {
			log.Printf(" %d - %v \n", i, err)
		}
		for _, ti := range t.Indizes {
			for idx, tif := range ti.Fields {
				stmt = Statement{
					Table:   tabDatTableIndex.Table,
					Command: INSERT,
					Presentation: Columns{
						{Name: "TABLE_NAME", Value: &t.Name},
						{Name: "NAME", Value: &ti.Name},
						{Name: "POS", Value: &idx},
						{Name: "PK", Value: &ti.PK},
						{Name: "UNIQ_IDX", Value: &ti.Unique},
						{Name: "FIELD", Value: &tif},
					},
				}
				i, err = ts.dialect.Exec(ts.connection, &stmt)
				if err != nil {
					log.Printf(" %d - %v \n", i, err)
				}

			}
		}
		if skip := sqlTableSkip[t.SQLName]; !skip && t.TableType == buts.DbTable {
			ts.createTableAndIndizes(t)
		}
	}
	return nil
}

func (ts *typeSystem) createTableAndIndizes(table buts.TableReg) {
	isPk := func(name string) bool {
		for _, i := range table.Indizes {
			if i.PK {
				for _, n := range i.Fields {
					if name == n {
						return true
					}
				}
			}
		}
		return false
	}
	if table.Kind == buts.Structure {
		tableStmt := Statement{
			Table:        table.SQLName,
			Command:      CREATE_TABLE,
			Presentation: make([]Column, 0),
		}
		if se, ok := ts.readStructure(table.Type); ok {
			for _, field := range se.Items {
				c := Column{
					Name: field.Name,
				}
				if field.Kind == buts.Element {
					if fe, ok := ts.readElement(field.Type); ok {
						var col ColumnDef
						switch fe.DbType {
						case buts.DbVarchar:
							col.Type = BT_VARCHAR
							col.Length = fe.DbLength
						case buts.DbInt:
							col.Type = BT_INT
						}
						col.PrimaryKey = isPk(field.Name)
						c.Value = col
					} else {
						panic("")
					}
				} else {
					panic("")
				}
				tableStmt.Presentation = append(tableStmt.Presentation, c)
			}
			if _, err := ts.dialect.Exec(ts.connection, &tableStmt); err != nil {
				panic(err)
			}
		}
	}
	/*

	 */
}

func (ts *typeSystem) readElement(name string) (te buts.ElementReg, ok bool) {
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
		return
	} else if rows.Next() {
		err = rows.Scan(
			&te.Name,
			&te.Description,
			&te.GoType,
			&te.DbType,
			&te.DbLength,
			&te.DbDecimals,
			&te.Tags,
			&te.Domain,
			&te.DomainTable,
			&te.DomainGoColumn,
			&te.DomainDbColumn,
			&te.Conversion,
		)
		if err != nil {
			log.Println(err)
			return
		} else {
			ok = true
			return
		}
	}
	return
}

func (ts *typeSystem) readStructure(name string) (te buts.StructureReg, ok bool) {
	stmt := Statement{
		Table:        tabDatStructure.Table,
		Command:      SELECT,
		Presentation: tabDatStructure.Presentation,
		Condition: Columns{
			{Name: "NAME", Value: &name, Operator: OP_EQ},
		},
	}
	rows, err := ts.dialect.Query(ts.connection, &stmt)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&te.Name, &te.Description, &te.Tags)
	} else {
		return
	}

	stmt = Statement{
		Table:        tabDatFeld.Table,
		Command:      SELECT,
		Presentation: Columns{{Name: "NAME"}, {Name: "DESCRIPTION"}, {Name: "KIND"}, {Name: "TYPE"}},
		Condition:    Columns{{Name: "STRUCT_NAME", Value: &name, Operator: OP_EQ}},
		Sort:         Columns{{Name: "POS", Value: "asc"}},
	}
	rows, err = ts.dialect.Query(ts.connection, &stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	te.Items = make([]buts.FeldReg, 0)
	for rows.Next() {
		var field buts.FeldReg
		err = rows.Scan(&field.Name, &field.Description, &field.Kind, &field.Type)
		if err != nil {
			log.Println(err)
			return
		}
		te.Items = append(te.Items, field)
	}
	ok = true
	return
}
