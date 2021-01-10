package sqldb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type sqlError struct {
	name   string
	action string
	column string
	err    string
}

func (e sqlError) Error() string {
	return fmt.Sprintf("Error %s of %s on %s: %s", e.action, e.name, e.column, e.err)
}

func CreateTable(db *sql.DB, stmt *Statement) (int, error) {
	b := new(strings.Builder)
	b.WriteString("CREATE TABLE ")
	b.WriteString(stmt.Table)
	b.WriteString(" (\n")
	var hasPK bool
	for i, col := range stmt.Presentation {
		if i > 0 {
			b.WriteString(",\n")
		}
		b.WriteString(col.Name)
		b.WriteString(" ")
		if ct, ok := col.Value.(ColumnDef); ok {
			hasPK = hasPK || ct.PrimaryKey
			switch ct.Type {
			case "VARCHAR":
				b.WriteString(fmt.Sprintf("VARCHAR(%d)", ct.Length))
			case "INT":
				b.WriteString("INT")
			}

		} else {
			return 0, sqlError{
				name:   stmt.Table,
				column: col.Name,
				action: "Create Table",
				err:    "col.Value isn't ColumnDef",
			}
		}
	}
	if hasPK {
		b.WriteString(",\n PRIMARY KEY ")
		b.WriteString(" (")
		var i int
		for _, col := range stmt.Presentation {
			ct := col.Value.(ColumnDef)
			if ct.PrimaryKey {
				if i > 0 {
					b.WriteString(",")
				}
				b.WriteString(col.Name)
				i++
			}
		}
		b.WriteString(")")

	}
	b.WriteString("\n)")
	stmtStr := b.String()
	log.Println(stmtStr)
	_, err := db.Exec(stmtStr)
	return 0, err
}
