package main

func StructFunc() {

	type Person struct {
		name string
		age  int
	}
	person := Person{"rishi", 20}

	println("person", person.name, "age", person.age)
}
