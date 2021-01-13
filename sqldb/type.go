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

// Field returns a struct type's i'th field type.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t typeNil) Field(i int) buts.Type {
	panic("no struct type")
}

// FieldByName returns the struct field type with the given name
// and a boolean indicating if the field was found.
func (t typeNil) FieldByName(name string) (buts.Type, bool, int) {
	panic("no struct type")
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (t typeNil) NumField() int {
	panic("no struct type")
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

// Field returns a struct type's i'th field type.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t *typeStructure) Field(i int) buts.Type {
	if 0 <= i && i < len(t.fields) {
		return t.fields[i].fieldType
	} else {
		panic("no struct type")
	}
}

// FieldByName returns the struct field type with the given name
// and a boolean indicating if the field was found.
func (t *typeStructure) FieldByName(name string) (buts.Type, bool, int) {
	for idx, field := range t.fields {
		if field.Name == name {
			return field.fieldType, true, idx
		}
	}
	return nil, false, -1
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (t *typeStructure) NumField() int {
	return len(t.fields)
}
