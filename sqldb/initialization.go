package sqldb

import (
	"reflect"

	"github.com/sblgo/buts"
)

var (
	defString   string
	defInt      int
	rtypeString = reflect.TypeOf(defString)
	rtypeInt    = reflect.TypeOf(defInt)
	goTypeMap   = map[buts.GoType]reflect.Type{
		buts.GoString: rtypeString,
		buts.GoInt:    rtypeInt,
	}
	dbTypeMap = map[buts.DbType]reflect.Type{
		buts.DbVarchar: rtypeString,
		buts.DbInt:     rtypeInt,
	}

	sqlTableSkip = map[string]bool{
		tabDatElement.Table:    true,
		tabDatStructure.Table:  true,
		tabDatFeld.Table:       true,
		tabDatTable.Table:      true,
		tabDatTableIndex.Table: true,
	}
)
