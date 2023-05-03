package validator

import (
	"fmt"
	"reflect"
)

type ValidatorRequired struct {
	Rule string
}

func (v *ValidatorRequired) IsValid(field string, data *map[string]any) bool {
	if _, ok := (*data)[field]; !ok {
		return false
	}
	value := (*data)[field]
	if reflect.TypeOf(value).String() == "string" {
		return len(value.(string)) > 0
	}

	return true
}

func (v *ValidatorRequired) GetErrorMessage(field string) string {
	return fmt.Sprintf("%s is required", field)
}
