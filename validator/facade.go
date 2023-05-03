package validator

import "strings"

type ValidatorInterface interface {
	IsValid(field string, data *map[string]any) bool
	GetErrorMessage(field string) string
}

type ValidatorFacade struct {
}

func (v *ValidatorFacade) GetValiator(rule string) ValidatorInterface {
	parseRule := rule
	if strings.Contains(rule, ":") {
		parseRule = rule[0:strings.Index(rule, ":")]
	}
	switch parseRule {
	case "required":
		return &ValidatorRequired{
			Rule: rule,
		}
	case "max":
		return &ValidatorMax{
			Rule: rule,
		}
	case "min":
		return &ValidatorMin{
			Rule: rule,
		}
	case "numeric":
		return &ValidatorNumeric{
			Rule: rule,
		}
	case "email":
		return &ValidatorEmail{
			Rule: rule,
		}
	case "phone":
		return &ValidatorPhone{
			Rule: rule,
		}
	}
	return nil
}
