package validator

import (
	"fmt"
	"regexp"

	"github.com/2751997nam/go-helpers/utils"
)

type ValidatorPhone struct {
	Rule string
}

func (v *ValidatorPhone) IsValid(field string, data *map[string]any) bool {
	if value, ok := (*data)[field]; ok {
		phone := utils.RegexReplace(`\D+`, "", value.(string))
		rex := regexp.MustCompile("^[0-9]{6,20}$")
		return !rex.MatchString(phone)
	}

	return true
}

func (v *ValidatorPhone) GetErrorMessage(field string) string {
	return fmt.Sprintf("%s is must be a number", field)
}
