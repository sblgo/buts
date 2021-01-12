package sqldb

import (
	"reflect"

	"github.com/sblgo/buts"
)

type typeNil struct {
	typeSystem  *typeSystem
	kind        buts.Kind
	name        string
	description string
	reflGoType  reflect.Type
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

func (t typeNil) Name() string {
	return t.name
}

func (t typeNil) TypeSystem() buts.TypeSystem {
	return t.typeSystem
}

func (t typeNil) ReflGoType() reflect.Type {
	return t.reflGoType
}

func (t typeNil) ReflDbType() reflect.Type {
	panic("not defined")
}

type typeElement struct {
	typeNil
	goType         buts.GoType
	dbType         buts.DbType
	dbLength       int
	dbDecimals     int
	tags           string
	domain         string
	domainTable    string
	domainGoColumn string
	domainDbColumn string
	conversion     string
	reflDbType     reflect.Type
}

func (t typeElement) GoType() buts.GoType {
	return t.goType
}
func (t typeElement) DbType() buts.DbType {
	return t.dbType
}

func (t typeElement) ReflDbType() reflect.Type {
	return t.reflDbType
}

type typeStructure struct {
	typeNil
	Description string
	Tags        string
	fields      []typeField
	reflFields  []reflect.StructField
}

type typeField struct {
	Name        string
	Description string
	Kind        buts.Kind
	Type        string
	fieldType   buts.Type
}

func (ts *typeStructure) GoType() buts.GoType {
	return buts.GoStructure
}
