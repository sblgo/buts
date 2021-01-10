package sqldb

import (
	"reflect"
)

var (
	tabDatElement = Statement{
		Table:   "TS_DAT_ELEMENT",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "COL_LENGTH", Value: ColumnDef{Type: BT_INT}},
			{Name: "COL_FRAC", Value: ColumnDef{Type: BT_INT}},
			{Name: "DB_TYPE", Value: ColumnDef{Type: BT_INT}},
			{Name: "GO_TYPE", Value: ColumnDef{Type: BT_INT}},
			{Name: "DESCRIPTION", Value: ColumnDef{Type: BT_VARCHAR, Length: 256}},
		},
	}
)

var (
	tableDefinitions = []Statement{
		tabDatElement,
	}
)

var (
	insDatElements = []struct {
		Name        string
		ColLength   int
		ColFrac     int
		DbType      int
		GoType      int
		Description string
	}{
		{"TDENAME", 20, 0, 1, int(reflect.String), "column name in ts_dat_element"},
	}
)
