package main

import (
	gsql "database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
	"github.com/sblgo/buts"
	"github.com/sblgo/buts/sqldb"
)

type v struct {
	msg string
}

func (v *v) Set(m string) {
	v.msg = m
}

func (v *v) Get() string {
	return v.msg
}

type V struct {
	*v
}

func New() V {
	return V{
		&v{},
	}
}

func change(v V) {
	x := v.Get()
	v.Set(x + x)
}

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
	v := New()
	v.Set("abc")
	change(v)
	fmt.Println(v.Get())
}
