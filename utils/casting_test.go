package utils

import "testing"

func Test_Casting(t *testing.T) {
	data := map[string]any{
		"id":         12312,
		"price":      "24.5",
		"high_price": 242.22232423,
	}
	type Tmp struct {
		ID        uint64
		Price     float32
		HighPrice uint64
	}
	tmp := Tmp{}
	MapToStruct(data, &tmp)
	Log(tmp)
}
