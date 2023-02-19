package common

import (
	"reflect"
	"unsafe"
)

func ReflectField(v interface{}, searchFieldName string) reflect.Value {

	var fieldValue reflect.Value

	t := reflect.TypeOf(v)
	o := reflect.ValueOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		o = o.Elem()
	}

	num := t.NumField()
	for i := 0; i < num; i++ {
		value := o.Field(i)
		fieldName := t.Field(i).Name
		if fieldName == searchFieldName {
			fieldValue = value
			break
		}
	}

	fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr()))
	return fieldValue
}
