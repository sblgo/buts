package buts

import "reflect"

// Type is an abstraction of a go type with additional features.
type Type interface {
	GoType() reflect.Type
}

func New(t Type) Value {

	return Value{
		valueType: t,
		//		reflectValue: reflect.New(t.ReflectType()),
	}
}
