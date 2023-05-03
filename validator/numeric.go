package validator

import (
	"fmt"

	"github.com/2751997nam/go-helpers/utils"
)

type ValidatorNumeric struct {
	Rule string
}

func (v *ValidatorNumeric) IsValid(field string, data *map[string]any) bool {
	if value, ok := (*data)[field]; ok {
		return utils.IsNumeric(value)
	}

	return true
}

func (v *ValidatorNumeric) GetErrorMessage(field string) string {
	return fmt.Sprintf("%s is must be a number", field)
}
