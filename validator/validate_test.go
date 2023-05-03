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
		"quantity": 20,
		"email":    "rockman2996@gmail.com",
		"phone":    "0977214760",
	}

	rules := map[string]string{
		"name":     "required|max:30",
		"price":    "required|numeric",
		"slug":     "min:3",
		"quantity": "required|min:10|max:20",
		"phone":    "required|phone",
		"email":    "required|email",
	}

	err := Validate(data, rules)
	if err != nil {
		log.Println(err.Error())
	}
}
