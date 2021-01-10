package sqldb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func SelectTable(db *sql.DB, stmt *Statement) (*sql.Rows, error) {
	b := new(strings.Builder)
	b.WriteString("SELECT ")
	for i, col := range stmt.Presentation {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(col.Name)
	}

	b.WriteString(" FROM ")
	b.WriteString(stmt.Table)
	conditions := make([]interface{}, 0)
	if len(stmt.Condition) > 0 {
		b.WriteString(" WHERE ")
		for i, col := range stmt.Condition {
			if i > 0 {
				b.WriteString(" AND ")
			}
			b.WriteString(fmt.Sprintf(" %s %s ? ", col.Name, col.Operator))
			conditions = append(conditions, col.Value)
		}
	}
	b.WriteString(" ;")
	query := b.String()
	log.Println(query)
	rows, err := db.Query(query, conditions...)
	return rows, err
}
