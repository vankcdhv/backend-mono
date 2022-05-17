package mysql

import (
	"backend-mono/cmd/utils"
	"fmt"
	"reflect"
	"time"
)

func GetListColumn(model interface{}, ignoreColumn []string, datetimeColumn []string) []string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	result := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("db")

		if jsonTag != "" && !utils.StringSliceContains(ignoreColumn, jsonTag) {
			if utils.StringSliceContains(datetimeColumn, jsonTag) {
				result = append(result, fmt.Sprintf("ROUND(UNIX_TIMESTAMP(%s)) as %s", jsonTag, jsonTag))
			} else {
				result = append(result, jsonTag)
			}
		}
	}
	return result
}

func GetListValues(model interface{}, ignoreColumn []string, datetimeColumn []string) []interface{} {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	result := make([]interface{}, 0)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		jsonTag := t.Field(i).Tag.Get("db")
		if jsonTag != "" && !utils.StringSliceContains(ignoreColumn, jsonTag) {
			if utils.StringSliceContains(datetimeColumn, jsonTag) {
				result = append(result, time.Unix(value.Int(), 0).Format("2006-01-02T15:04:05"))
			} else {
				result = append(result, value.Interface())
			}
		}
	}
	return result
}
