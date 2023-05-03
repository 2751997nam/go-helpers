package validator

import (
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	data := map[string]any{
		"name":     "asdsd",
		"price":    13.5,
		"slug":     "xxxxx",
		"quantity": 50,
	}

	rules := map[string]string{
		"name":     "required|max:30",
		"price":    "required|numeric",
		"slug":     "min:3",
		"quantity": "required|min:10|max:20",
	}

	err := Validate(data, rules)
	if err != nil {
		log.Println(err.Error())
	}
}
