package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/2751997nam/go-helpers/utils"
)

type ValidatorMin struct {
	Rule string
	Min  uint64
	Type string
}

func (v *ValidatorMin) IsValid(field string, data *map[string]any) bool {
	if value, ok := (*data)[field]; ok {
		v.Min = utils.AnyToUint(v.Rule[strings.Index(v.Rule, ":")+1:])
		if reflect.TypeOf(value).String() == "string" {
			v.Type = "string"
			return len(value.(string)) >= int(v.Min)
		}
		castedValue, err := utils.CastType(value, float64(0))
		if err == nil {
			v.Type = "number"
			return castedValue.Interface().(float64) >= float64(v.Min)
		}
		return true
	}

	return true
}

func (v *ValidatorMin) GetErrorMessage(field string) string {
	if v.Type == "string" {
		return fmt.Sprintf("field %s must have at least %d characters", field, v.Min)
	}
	return fmt.Sprintf("field %s must be greater or equal to %d", field, v.Min)
}
