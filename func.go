package main

import (
	f "fmt" // fmt module to f name
)

func init() { // first auto run
	f.Println("init")
}

func main() { // secound auto run
	f.Println("main func")
	f.Println(max(1, 2))

	var num_sum, num_times = sumAndTimes(1, 2)

	f.Println("num_sum: ", num_sum)
	f.Println("num_times: ", num_times)

	fncMultiArg(7, 8, 9, 10)

	x := 10
	f.Println("add1(x): ", add1(x))
	f.Println("add2(x): ", add2(&x)) // pass pointer

	f.Println("x", x)   // value
	f.Println("&x", &x) // address
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func sumAndTimes(a int, b int) (sum int, times int) {
	sum = a + b
	times = a * b
	return
}

func fncMultiArg(arg ...int) {
	for _, n := range arg {
		f.Printf("And the number is: %d\n", n)
	}
}

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int { // change address value
	*a = *a + 1
	return *a
}
