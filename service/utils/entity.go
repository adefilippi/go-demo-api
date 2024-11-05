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

func FiltersToWhereQuery(filters map[string]interface{}, model interface{}) (string, []interface{}) {

	if len(filters) == 0 {
		return "", nil
	}

	var whereClauses []string
	var args []interface{}

	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		filterTag := fieldType.Tag.Get("filter")
		if filterTag == "" {
			continue
		}
		filterValue, ok := filters[filterTag]
		if !ok {
			continue
		}
		columnName := strings.Split(jsonTag, ",")[0]
		if fieldType.Type.Kind() == reflect.String || (fieldType.Type.Kind() == reflect.Ptr && fieldType.Type.Elem().Kind() == reflect.String) {
			whereClauses = append(whereClauses, fmt.Sprintf("LOWER(%s) = LOWER(?)", columnName))
			args = append(args, filterValue)
		} else {
			whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", columnName))
			args = append(args, filterValue)
		}
	}

	// Join where clauses with " AND " and return
	whereClause := strings.Join(whereClauses, " AND ")
	return whereClause, args
}
