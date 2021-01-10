package sqldb

import (
	"reflect"

	"github.com/sblgo/buts"
)

type typeNil struct {
	typeSystem *typeSystem
	kind       buts.Kind
}

func (t typeNil) DbType() buts.DbType {
	panic("not defined")
}

func (t typeNil) GoType() buts.GoType {
	panic("not defined")
}

func (t typeNil) Kind() buts.Kind {
	return t.kind
}

func (t typeNil) TypeSystem() buts.TypeSystem {
	return t.typeSystem
}

type typeElement struct {
	typeNil
	buts.ElementReg
	refGoType reflect.Type
	refDbType reflect.Type
}

func (t typeElement) GoType() buts.GoType {
	return t.ElementReg.GoType
}
func (t typeElement) DbType() buts.DbType {
	return t.ElementReg.DbType
}

func (te *typeElement) New() buts.Value {
	return buts.Value{}
}
