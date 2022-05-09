package mysql

import (
	"reflect"
)

func GetListColumn(model interface{}) []string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	result := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag != "" {
			result[i] = jsonTag
		}
	}
	return result
}

func GetListValues(model interface{}) []interface{} {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		result[i] = value.Interface()
	}
	return result
}
