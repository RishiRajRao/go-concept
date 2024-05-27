package main

import "fmt"

// struct type
type PersonName struct {
	firstName, lastName string
}

func (person PersonName) MethodsFunc(title string) string {
	return fmt.Sprintf(person.firstName + " " + person.lastName + " " + title)
}

// non-struct type
type MyInt int

func (num MyInt) Sum(val int) int {
	return int(num) * val
}

//pointer struct & methods

type Address struct {
	city, pincode string
}

func (address *Address) FullAddress() {
	address.city = "Patna"
	address.pincode = "201301"
}
