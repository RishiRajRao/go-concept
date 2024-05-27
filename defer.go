package main

func DeferApply(temp int) {
	// temp := 10
	defer func(temp int) {
		println("temp is ", temp)
	}(temp)

	temp += 1

	println("updated temp is", temp)
}

func MultiDefer() {

	for _, v := range []int{1, 2, 3} {
		defer println("value is", v)
	}

	println("out of stack")

}

// defer println("Hey first defer in main")
// MultiDefer()
// defer println("Hey last defer in main")
// In main stack they have their own order while child/nested function has their own order of stack.
