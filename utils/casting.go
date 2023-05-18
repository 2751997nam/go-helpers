package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func AnyToString(value any) string {
	if value != nil {
		return strings.Trim(fmt.Sprint(value), " ")
	}

	return ""
}

func AnyToInt(value any) int {
	result, _ := strconv.Atoi(AnyToString(value))
	return result
}

func AnyToUint(value any) uint64 {
	result, _ := strconv.ParseInt(AnyToString(value), 10, 64)
	return uint64(result)
}

func AnyFloat64ToUint64(value any) uint64 {
	var result float64 = 0
	if reflect.TypeOf(value).Name() == "string" {
		result, _ = strconv.ParseFloat(AnyToString(value), 64)
	} else {
		result = value.(float64)
	}

	return uint64(result)
}

func AnyToFloat(value any) float32 {
	result, _ := strconv.ParseFloat(AnyToString(value), 64)
	return float32(result)
}

func CastType(a any, b any) (reflect.Value, error) {
	// Get the type of variable b
	bType := reflect.ValueOf(b).Type()

	// Create a new variable of the same type as b
	newB := reflect.New(bType).Elem()

	// Cast variable a to the type of variable b using a type assertion
	castedA := reflect.ValueOf(a).Convert(bType)

	// Check if the types of a and b are compatible
	if !castedA.Type().AssignableTo(bType) {
		return reflect.Value{}, fmt.Errorf("a cannot be cast to the type of b")
	}

	// Assign the casted value of a to the new variable of type b
	newB.Set(castedA)

	// Return the new variable of type b
	return newB, nil
}

func MapToStruct[T any](mapValue map[string]any, structValue *T) {
	for key, value := range mapValue {
		rex := regexp.MustCompile(`(\b|_)(\w)`)
		field := rex.ReplaceAllStringFunc(key, func(str string) string {
			if len(str) > 0 {
				return strings.ToUpper(string(str[len(str)-1]))
			}
			return ""
		})

		if strings.ToLower(field) == "id" {
			field = "ID"
		}

		structField := reflect.ValueOf(structValue).Elem().FieldByName(field)
		if structField.IsValid() {
			val, _ := CastType(value, structField.Interface())
			structField.Set(val)
		}
	}
}
