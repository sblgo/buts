package main

import (
	gsql "database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
	"github.com/sblgo/buts"
	"github.com/sblgo/buts/sqldb"
)

func main() {
	db, err := gsql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var ts buts.TypeSystem
	ts, err = sqldb.New("standard", db)
	fmt.Printf("%v - %v\n", ts, err)
	t := ts.New(buts.Element, "TDENAME")
	fmt.Printf("%v\n", t)
	v := buts.New(t)
	fmt.Printf("%v\n", v)
	v.SetInt(3)
	fmt.Println(v.Int())
	v.SetString("ABCDEF" + v.String())
	fmt.Println(v.String())
}
