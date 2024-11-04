package utils

import (
	"strings"
	"reflect"
	"fmt"
)

func GetAssociationValueId(data interface{}, tag string) string {
	tag = tag + "_id"
	val := reflect.ValueOf(data)

	// Dereference the pointer if necessary
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if strings.Contains(field.Tag.Get("json"), tag) {
			result := val.Field(i).Interface()
			return fmt.Sprintf("%s", result)
		}
	}

	return ""
}
