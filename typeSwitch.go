package main

func TypeSwitch(val interface{}) {
	switch val.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case float64:
		println("float64")
	case bool:
		println("bool")
	default:
		println("Out of list")
	}
}
