package validator

import (
	"fmt"
	"strings"
)

func Validate(data map[string]any, rules map[string]string) error {
	for field, rule := range rules {
		ruleItems := strings.Split(rule, "|")
		for _, ruleItem := range ruleItems {
			builder := ValidatorFacade{}
			validator := builder.GetValiator(ruleItem)
			if validator == nil {
				return fmt.Errorf("no validator for %s rule", ruleItem)
			} else if !validator.IsValid(field, &data) {
				return fmt.Errorf(validator.GetErrorMessage(field))
			}
		}
	}
	return nil
}
