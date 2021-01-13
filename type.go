package buts

import (
	"reflect"
)

//
// Kinds of ConversionFunc
// - conv_db_<type>_read
// - conv_db_<type>_write
// - conv_ex_<type>_in
// - conv_ex_<type>_out
type ConversionFunc func(desc Type, k reflect.Kind, v interface{}) (interface{}, error)

// Type is an abstraction of a go type with additional features.
type Type interface {
	GoType() GoType
	DbType() DbType
	Kind() Kind
	ReflGoType() reflect.Type
	ReflDbType() reflect.Type

	//TypeSystem returns the  type system of the type
	TypeSystem() TypeSystem

	// Name returns the type's name.
	Name() string

	// Field returns a struct type's i'th field type.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int) Type

	// FieldByName returns the struct field type with the given name
	// ,a boolean indicating if the field was found and the index of the field
	// in the structure.
	FieldByName(name string) (Type, bool, int)

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int
}

func New(t Type) Value {
	v := Value{
		&value{
			valueType:    t,
			reflectValue: reflect.New(t.ReflGoType()),
		},
	}

	return v
}
