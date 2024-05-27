package main

func Looping(number []int) {

	sum := 0
	for _, val := range number {
		sum += val
	}

	println("sum loop is", sum)
}

func RawLoop(number int) {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	println("sum loop is raw", sum)
}
