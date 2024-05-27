package main

import "fmt"

func MapFunc() {
	myMap := map[string]string{"first": "rishi", "last": "raj"}

	// myMap := make(map[string]string)
	myMap["final"] = "Rao"
	// myMap["last"] = "Rao"

	delete(myMap, "final")

	if val, ok := myMap["first"]; ok {
		fmt.Println("hey key is final", val)
	} else {
		fmt.Println("my map is", myMap, val, ok)
	}

	fmt.Println("my map is", myMap)

}
