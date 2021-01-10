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
	TypeSystem() TypeSystem
}

func New(t Type) Value {
	v := Value{
		&value{
			valueType: t,
		},
	}
	switch t.Kind() {
	case Element:
		switch t.GoType() {
		case GoString:
			var s string
			v.reflectValue = reflect.New(reflect.TypeOf(s))
		case GoInt:
			var i int
			v.reflectValue = reflect.New(reflect.TypeOf(i))
		}
	}

	return v
}
