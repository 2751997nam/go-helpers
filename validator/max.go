package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/2751997nam/go-helpers/utils"
)

type ValidatorMax struct {
	Rule string
	Max  uint64
	Type string
}

func (v *ValidatorMax) IsValid(field string, data *map[string]any) bool {
	if value, ok := (*data)[field]; ok {
		v.Max = utils.AnyToUint(v.Rule[strings.Index(v.Rule, ":")+1:])
		if reflect.TypeOf(value).String() == "string" {
			v.Type = "string"
			return len(value.(string)) <= int(v.Max)
		}
		castedValue, err := utils.CastType(value, float64(0))
		if err == nil {
			v.Type = "number"
			return castedValue.Interface().(float64) <= float64(v.Max)
		}
		return true
	}

	return true
}

func (v *ValidatorMax) GetErrorMessage(field string) string {
	if v.Type == "string" {
		return fmt.Sprintf("field %s must be no more than %d characters", field, v.Max)
	}
	return fmt.Sprintf("field %s must not greater than %d", field, v.Max)
}
