package sqldb

import (
	"github.com/sblgo/buts"
)

var (
	tabDatElement = Statement{
		Table:   "TS_DAT_ELEMENT",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "DESCRIPTION", Value: ColumnDef{Type: BT_VARCHAR, Length: 256}},
			{Name: "GO_TYPE", Value: ColumnDef{Type: BT_INT}},
			{Name: "DB_TYPE", Value: ColumnDef{Type: BT_INT}},
			{Name: "COL_LENGTH", Value: ColumnDef{Type: BT_INT}},
			{Name: "COL_FRAC", Value: ColumnDef{Type: BT_INT}},
			{Name: "TAGS", Value: ColumnDef{Type: BT_VARCHAR, Length: 2048}},
			{Name: "DOMAIN_NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 48}},
			{Name: "DOMAIN_TABLE", Value: ColumnDef{Type: BT_VARCHAR, Length: 48}},
			{Name: "DOMAIN_GOCOL", Value: ColumnDef{Type: BT_VARCHAR, Length: 48}},
			{Name: "DOMAIN_DBCOL", Value: ColumnDef{Type: BT_VARCHAR, Length: 48}},
			{Name: "DOMAIN_CONV", Value: ColumnDef{Type: BT_VARCHAR, Length: 100}},
		},
	}
)

var (
	tableDefinitions = []Statement{
		tabDatElement,
	}
)

var (
	insDatElements = []buts.ElementReg{
		{"TDENAME", "column name in ts_dat_element", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDEDESC", "column description in ts_dat_element", buts.GoString, buts.DbVarchar, 256, 0, "", "", "", "", "", ""},
		{"TDEGOTYP", "column go_type in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEDBTYP", "column db_type in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDECLLEN", "column col_length in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDECLFRC", "column col_frack in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEDNAME", "column domain_name in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEDTABL", "column domain_table in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEDGOCL", "column domain_gocol in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEDCONV", "column domain_conv in ts_dat_element", buts.GoString, buts.DbVarchar, 100, 0, "", "", "", "", "", ""},
	}
)
