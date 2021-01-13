package buts

import (
	"reflect"
)

type Value struct {
	*value
}

type value struct {
	valueType    Type
	reflectValue reflect.Value
	exin, exout  ConversionFunc
}

func (v *value) convFuncSet() {
	if v.exin == nil {
		v.exin = convFunc("Ex", v.valueType.GoType().String(), "In")
	}
}

func (v *value) convFuncGet() {
	if v.exout == nil {
		v.exout = convFunc("Ex", v.valueType.GoType().String(), "Out")
	}
}

func (v *value) set(k reflect.Kind, i interface{}) {
	v.convFuncSet()
	p, err := v.exin(v.valueType, k, i)
	if err != nil {
		panic(err)
	}
	v.reflectValue.Elem().Set(reflect.ValueOf(p))
}

func (v *value) get(k reflect.Kind) reflect.Value {
	v.convFuncGet()
	iv := v.reflectValue
	if iv.Kind() == reflect.Ptr {
		iv = iv.Elem()
	}
	p, err := v.exout(v.valueType, k, iv.Interface())
	if err != nil {
		panic(err)
	}
	pv := reflect.ValueOf(p)
	if pv.Kind() == reflect.Ptr {
		pv = pv.Elem()
	}
	return pv
}

func (v *value) SetString(s string) {
	v.set(reflect.String, s)
}

func (v *value) SetInt(i int) {
	v.set(reflect.Int, i)
}

func (v *value) String() string {
	return v.get(reflect.String).Interface().(string)
}

func (v *value) Int() int {
	return v.get(reflect.Int).Interface().(int)
}

func (v *value) Type() Type {
	return v.valueType
}

// Field returns a struct type's i'th field type.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t *value) Field(i int) Value {
	if 0 <= i && i < t.Type().NumField() {
		return Value{
			value: &value{
				valueType:    t.Type().Field(i),
				reflectValue: t.reflectValue.Elem().Field(i).Addr(),
			},
		}
	} else {
		panic("no struct type")
	}
}

// FieldByName returns the struct field type with the given name
// and a boolean indicating if the field was found.
func (t *value) FieldByName(name string) (Value, bool) {
	if ft, ok, idx := t.Type().FieldByName(name); ok {
		return Value{
			&value{
				valueType:    ft,
				reflectValue: t.reflectValue.Elem().Field(idx).Addr(),
			},
		}, true
	}
	return Value{}, false
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (v *value) NumField() int {
	return v.Type().NumField()
}
