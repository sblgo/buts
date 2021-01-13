package main

import (
	"database/sql"
	"log"

	_ "github.com/proullon/ramsql/driver"
	"github.com/sblgo/buts"

	"github.com/sblgo/buts/sqldb"
)

var (
	regElement = []buts.ElementReg{
		{"TSTELM1", "", buts.GoString, buts.DbVarchar, 32, 0, "", "", "", "", "", ""},
	}
	regStructure = []buts.StructureReg{
		{"TST_STRUCT", "", "", []buts.FeldReg{
			{"TSTFIEL1", "", buts.Element, "TSTELM1"},
		}},
	}
	regTable = []buts.TableReg{
		{"TST_TABLE", "", "", buts.DbTable, buts.Structure, "TST_STRUCT", "DB_TST_TABLE", []buts.TableIndex{
			{"PK", true, true, []string{"TSTFIEL1"}},
		}},
	}
)

func main() {
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		log.Fatal(err)
	}
	ts, err := sqldb.New("standard", db)
	if err != nil {
		log.Fatal(err)
	}

	ts.Register(regElement, regStructure, regTable)
}
