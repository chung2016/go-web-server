package main

import "fmt"

func main() {
	fmt.Println("func")
	fmt.Println(max(1, 2))

	var num_sum, num_times = sumAndTimes(1, 2)

	fmt.Println("num_sum: ", num_sum)
	fmt.Println("num_times: ", num_times)

	fncMultiArg(7, 8, 9, 10)

	x := 10
	fmt.Println("add1(x): ", add1(x))
	fmt.Println("add2(x): ", add2(&x)) // pass pointer

	fmt.Println("x", x)   // value
	fmt.Println("&x", &x) // address
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
		fmt.Printf("And the number is: %d\n", n)
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
