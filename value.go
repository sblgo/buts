package buts

import "reflect"

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
