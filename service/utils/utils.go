package utils

import (
	"database/sql"
	"reflect"
	"github.com/google/uuid"
	"fmt"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ScanStruct(dest interface{}, rows *sql.Rows) error {
	v := reflect.ValueOf(dest).Elem()
	t := v.Type()

	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		values[i] = v.Field(i).Addr().Interface()
	}

	if err := rows.Scan(values...); err != nil {
		return err
	}

	return nil
}

func ParseId(id interface{}) (uuid.UUID, error) {
	idStr, ok := id.(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("expected string for id, got %T", id)
	}

	uid, error := uuid.Parse(idStr)
	if error != nil {
		return uuid.UUID{}, error
	}

	return uid, nil
}
