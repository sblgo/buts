package sqldb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func InsertTable(db *sql.DB, stmt *Statement) (int, error) {
	b := new(strings.Builder)
	b.WriteString("INSERT INTO ")
	b.WriteString(stmt.Table)
	b.WriteString(" (")
	for i, col := range stmt.Presentation {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(col.Name)
	}

	b.WriteString(" ) VALUES ( ")
	conditions := make([]interface{}, 0)
	for i, col := range stmt.Presentation {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("? "))
		conditions = append(conditions, col.Value)
	}
	b.WriteString(")")
	query := b.String()
	result, err := db.Exec(query, conditions...)
	if err != nil {
		log.Println(query)
		return 0, err
	}
	rows, err := result.RowsAffected()
	return int(rows), err
}
