package sarif

import (
	"reflect"
	"strings"
)

func initializeArrays(v interface{}) {
	initializeArraysValue(reflect.ValueOf(v))
}

func initializeArraysValue(val reflect.Value) {
	initializeArraysValueWithParent(val, reflect.StructField{})
}

func initializeArraysValueWithParent(val reflect.Value, structField reflect.StructField) {
	if !val.IsValid() {
		return
	}

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return
		}
		initializeArraysValueWithParent(val.Elem(), structField)

	case reflect.Interface:
		if val.IsNil() {
			return
		}
		initializeArraysValueWithParent(val.Elem(), structField)

	case reflect.Slice:
		if val.IsNil() && val.CanSet() {
			if !hasOmitEmpty(structField) {
				val.Set(reflect.MakeSlice(val.Type(), 0, 0))
			}
		}
		for i := 0; i < val.Len(); i++ {
			initializeArraysValueWithParent(val.Index(i), reflect.StructField{})
		}

	case reflect.Struct:
		structType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := structType.Field(i)
			if field.CanSet() || field.CanInterface() {
				initializeArraysValueWithParent(field, fieldType)
			}
		}

	case reflect.Map:
		if val.IsNil() {
			return
		}
		for _, key := range val.MapKeys() {
			initializeArraysValueWithParent(val.MapIndex(key), reflect.StructField{})
		}
	}
}

func hasOmitEmpty(field reflect.StructField) bool {
	tag := field.Tag.Get("json")
	if tag == "" {
		return false
	}
	for i, part := range strings.Split(tag, ",") {
		if i > 0 && part == "omitempty" {
			return true
		}
	}
	return false
}
