package helper

import (
	"go-exercise/customer-manage/model"
)

//TypeInference 类型推断
func TypeInference(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	case model.Customer:
		return "customer"
	default:
		return "defaullt"
	}
}
