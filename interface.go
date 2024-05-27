package main

import "fmt"

var human Human

type Man struct {
	name string
	age  int
}

func (man *Man) Identity() string {
	return fmt.Sprintf("Age is %v and name is %s", man.age, man.name)
}

//interface definition

type Human interface {
	Identity() string
}

func InterfaceFunc() {
	man := Man{"Rishi", 30}
	human = &man
	println(human.Identity())
}

func InterfaceGenric() {
	var hero interface{} = "Rishi"

	intHero, ok := hero.(int)
	fmt.Printf("my value is %d  and %t \n", intHero, ok)

}
