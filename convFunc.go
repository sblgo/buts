package buts

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	errorMissingConvRule = errors.New("Missing conversion rule")
	convFuncMap          = map[string]ConversionFunc{
		"ConvExGoStringIn":     ConvExGoStringIn,
		"ConvExGoStringOut":    ConvExGoStringOut,
		"ConvExGoIntIn":        ConvExGoIntIn,
		"ConvExGoIntOut":       ConvExGoIntOut,
		"ConvExGoStructureIn":  ConvExGoStructureIn,
		"ConvExGoStructureOut": ConvExGoStructureOut,
	}
)

func unwrapReflect(v interface{}) interface{} {
	if r, ok := v.(reflect.Value); ok {
		if r.Kind() == reflect.Ptr {
			return r.Elem().Interface()
		} else {
			return r.Interface()
		}
	} else {
		return v
	}
}

func ConvExGoStringIn(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	u := unwrapReflect(v)
	switch t := u.(type) {
	case string:
		return t, nil
	case int:
		return strconv.Itoa(t), nil
	}
	return nil, errorMissingConvRule
}

func ConvExGoStringOut(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	u := unwrapReflect(v).(string)
	switch k {
	case reflect.String:
		return u, nil
	case reflect.Int:
		if i, err := strconv.Atoi(u); err != nil {
			return nil, err
		} else {
			return i, nil
		}
	}
	return nil, errorMissingConvRule
}

func ConvExGoIntIn(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	u := unwrapReflect(v)
	switch t := u.(type) {
	case string:
		if i, err := strconv.Atoi(t); err != nil {
			return nil, err
		} else {
			return i, nil
		}
	case int:
		return t, nil
	}
	return nil, errorMissingConvRule
}

func ConvExGoIntOut(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	u := unwrapReflect(v).(int)
	switch k {
	case reflect.String:
		return strconv.Itoa(u), nil
	case reflect.Int:
		return u, nil
	}
	return nil, errorMissingConvRule
}

func ConvExGoStructureIn(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	return nil, errorMissingConvRule
}

func ConvExGoStructureOut(desc Type, k reflect.Kind, v interface{}) (interface{}, error) {
	u := unwrapReflect(v)
	switch k {
	case reflect.String:
		return fmt.Sprintf("%v", u), nil
	case reflect.Int:
		return 0, errorMissingConvRule
	}
	return desc.Name(), nil
}

func convFunc(e string, t string, d string) ConversionFunc {
	name := "Conv" + e + strings.ToUpper(t[0:1]) + t[1:] + d
	if f, ok := convFuncMap[name]; ok {
		return f
	}
	return nil
}
