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
	tabDatStructure = Statement{
		Table:   "TS_DAT_STRUCTURE",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "DESCRIPTION", Value: ColumnDef{Type: BT_VARCHAR, Length: 256}},
			{Name: "TAGS", Value: ColumnDef{Type: BT_VARCHAR, Length: 2048}},
		},
	}
	tabDatFeld = Statement{
		Table:   "TS_DAT_FELD",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "STRUCT_NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "POS", Value: ColumnDef{Type: BT_INT, PrimaryKey: true}},
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20}},
			{Name: "DESCRIPTION", Value: ColumnDef{Type: BT_VARCHAR, Length: 256}},
			{Name: "KIND", Value: ColumnDef{Type: BT_INT}},
			{Name: "TYPE", Value: ColumnDef{Type: BT_VARCHAR, Length: 20}},
		},
	}
	tabDatTable = Statement{
		Table:   "TS_DAT_TABLE",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "DESCRIPTION", Value: ColumnDef{Type: BT_VARCHAR, Length: 256}},
			{Name: "TAGS", Value: ColumnDef{Type: BT_VARCHAR, Length: 2048}},
			{Name: "TABLE_TYPE", Value: ColumnDef{Type: BT_INT}},
			{Name: "KIND", Value: ColumnDef{Type: BT_INT}},
			{Name: "TYPE", Value: ColumnDef{Type: BT_VARCHAR, Length: 20}},
			{Name: "SQL_NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20}},
		},
	}
	tabDatTableIndex = Statement{
		Table:   "TS_DAT_TABIDX",
		Command: CREATE_TABLE,
		Presentation: []Column{
			{Name: "TABLE_NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "NAME", Value: ColumnDef{Type: BT_VARCHAR, Length: 20, PrimaryKey: true}},
			{Name: "POS", Value: ColumnDef{Type: BT_INT, PrimaryKey: true}},
			{Name: "PK", Value: ColumnDef{Type: BT_INT}},
			{Name: "UNIQ_IDX", Value: ColumnDef{Type: BT_INT}},
			{Name: "FIELD", Value: ColumnDef{Type: BT_VARCHAR, Length: 20}},
		},
	}
)

var (
	tableDefinitions = []Statement{
		tabDatElement,
		tabDatStructure,
		tabDatFeld,
		tabDatTable,
		tabDatTableIndex,
	}
)

var (
	insDatElements = []buts.ElementReg{
		// ts_dat_element
		{"TDEENAME", "column name in ts_dat_element", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDEEDESC", "column description in ts_dat_element", buts.GoString, buts.DbVarchar, 256, 0, "", "", "", "", "", ""},
		{"TDEEGOTYP", "column go_type in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEEDBTYP", "column db_type in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEETAG", "column tag in ts_dat_element", buts.GoString, buts.DbVarchar, 2024, 0, "", "", "", "", "", ""},
		{"TDEECLLEN", "column col_length in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEECLFRC", "column col_frack in ts_dat_element", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEEDNAME", "column domain_name in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEEDTABL", "column domain_table in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEEDGOCL", "column domain_gocol in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEEDDBCL", "column domain_dbcol in ts_dat_element", buts.GoString, buts.DbVarchar, 48, 0, "", "", "", "", "", ""},
		{"TDEEDCONV", "column domain_conv in ts_dat_element", buts.GoString, buts.DbVarchar, 100, 0, "", "", "", "", "", ""},
		// ts_dat_structure
		{"TDESNAME", "column name in ts_dat_structure", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDESDESC", "column description in ts_dat_structure", buts.GoString, buts.DbVarchar, 256, 0, "", "", "", "", "", ""},
		{"TDESTAG", "column tag in ts_dat_structure", buts.GoString, buts.DbVarchar, 2024, 0, "", "", "", "", "", ""},
		{"TDEFSNAME", "column struct_name in ts_dat_feld", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		// ts_dat_feld
		{"TDEFPOS", "column pos in ts_dat_feld", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEFNAME", "column name in ts_dat_feld", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDEFDESC", "column description in ts_dat_feld", buts.GoString, buts.DbVarchar, 256, 0, "", "", "", "", "", ""},
		{"TDEFKIND", "column kind in ts_dat_feld", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEFTYPE", "column type in ts_dat_feld", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		// TS_DAT_TABLE
		{"TDETNAME", "column name in ts_dat_table", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDETDESC", "column description in ts_dat_table", buts.GoString, buts.DbVarchar, 256, 0, "", "", "", "", "", ""},
		{"TDETTAG", "column tag in ts_dat_table", buts.GoString, buts.DbVarchar, 2024, 0, "", "", "", "", "", ""},
		{"TDETTABT", "column table_type in ts_dat_table", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDETKIND", "column kind in ts_dat_table", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDETTYPE", "column type in ts_dat_table", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDETSQLNM", "column name in ts_dat_table", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		// TS_DAT_TABIDX
		{"TDEITNAME", "column name in ts_dat_tabidx", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDEINAME", "column name in ts_dat_tabidx", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
		{"TDEIPOS", "column pos in ts_dat_tabidx", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEIPK", "column pos in ts_dat_tabidx", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEIUNIQ", "column pos in ts_dat_tabidx", buts.GoInt, buts.DbInt, 0, 0, "", "", "", "", "", ""},
		{"TDEIFELD", "column name in ts_dat_tabidx", buts.GoString, buts.DbVarchar, 20, 0, "", "", "", "", "", ""},
	}

	insDatStructure = []buts.StructureReg{
		{"TDSELEMENT", "table ts_dat_element", "", []buts.FeldReg{
			{"NAME", "NAME of element", buts.Element, "TDEENAME"},
			{"DESCRIPTION", "DESCRIPTION of element", buts.Element, "TDEEDESC"},
			{"GO_TYPE", "", buts.Element, "TDEEGOTYP"},
			{"DB_TYPE", "", buts.Element, "TDEEDBTYP"},
			{"COL_LENGTH", "", buts.Element, "TDEECLLEN"},
			{"COL_FRAC", "", buts.Element, "TDEECLFRC"},
			{"TAGS", "", buts.Element, "TDEETAG"},
			{"DOMAIN_NAME", "", buts.Element, "TDEEDNAME"},
			{"DOMAIN_TABLE", "", buts.Element, "TDEEDTABL"},
			{"DOMAIN_GOCOL", "", buts.Element, "TDEEDGOCL"},
			{"DOMAIN_DBCOL", "", buts.Element, "TDEEDDBCL"},
			{"DOMAIN_CONV", "", buts.Element, "TDEEDCONV"},
		}},
		{"TDSSTRUCT", "table ts_dat_structure", "", []buts.FeldReg{
			{"NAME", "", buts.Element, "TDESNAME"},
			{"DESCRIPTION", "", buts.Element, "TDESDESC"},
			{"TAGS", "", buts.Element, "TDESTAG"},
		}},
		{"TDSFELD", "table ts_dat_feld", "", []buts.FeldReg{
			{"STRUCT_NAME", "", buts.Element, "TDEFSNAME"},
			{"POS", "", buts.Element, "TDEFPOS"},
			{"NAME", "", buts.Element, "TDEFNAME"},
			{"DESCRIPTION", "", buts.Element, "TDEFDESC"},
			{"KIND", "", buts.Element, "TDEFKIND"},
			{"TYPE", "", buts.Element, "TDEFTYPE"},
		}},
		{"TDSTABLE", "table ts_dat_table", "", []buts.FeldReg{
			{"NAME", "", buts.Element, "TDETNAME"},
			{"DESCRIPTION", "", buts.Element, "TDETDESC"},
			{"TAGS", "", buts.Element, "TDETTAG"},
			{"TABLE_TYPE", "", buts.Element, "TDETTABT"},
			{"KIND", "", buts.Element, "TDETKIND"},
			{"TYPE", "", buts.Element, "TDETTYPE"},
			{"SQL_NAME", "", buts.Element, "TDETSQLNM"},
		}},
		{"TDSTABIDX", "table ts_dat_tabidx", "", []buts.FeldReg{
			{"TABLE_NAME", "", buts.Element, "TDEITNAME"},
			{"NAME", "", buts.Element, "TDEINAME"},
			{"POS", "", buts.Element, "TDEIPOS"},
			{"PK", "", buts.Element, "TDEIPK"},
			{"TABLE_TYPE", "", buts.Element, "TDEIUNIQ"},
			{"UNIQ_IDX", "", buts.Element, "TDEIUNIQ"},
			{"FIELD", "", buts.Element, "TDEIFELD"},
		}},
	}
	insDatTable = []buts.TableReg{
		{"TDTELEMENT", "Table for elements", "", buts.DbTable, buts.Structure, "TDSELEMENT", "TS_DAT_ELEMENT", []buts.TableIndex{
			{"PK", true, true, []string{"NAME"}},
		}},
		{"TDTSTRUCT", "Table for structures", "", buts.DbTable, buts.Structure, "TDSSTRUCT", "TS_DAT_STRUCTURE", []buts.TableIndex{
			{"PK", true, true, []string{"NAME"}},
		}},
		{"TDTFELD", "Table for structures", "", buts.DbTable, buts.Structure, "TDSFELD", "TS_DAT_FELD", []buts.TableIndex{
			{"PK", true, true, []string{"STRUCT_NAME", "POS"}},
		}},
		{"TDTTABLE", "Table for tables", "", buts.DbTable, buts.Structure, "TDSTABLE", "TS_DAT_TABLE", []buts.TableIndex{
			{"PK", true, true, []string{"NAME"}},
		}},
		{"TDTTABIDX", "Table for indizes", "", buts.DbTable, buts.Structure, "TDSTABIDX", "TS_DAT_STRUCTURE", []buts.TableIndex{
			{"PK", true, true, []string{"TABLE_NAME", "NAME", "POS"}},
		}},
	}
)
