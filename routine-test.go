package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // do front helf sum
	go sum(a[len(a)/2:], c) // do back helf sum
	x, y := <-c, <-c        // receive from c

	z := make(chan int)
	go sum(a, z)
	zz := <-z
	fmt.Println("zz sum: ", zz)
	fmt.Println("x: ", x, "y: ", y, "sum: ", x+y)
}
