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
	calcSum()
	chansize()

	runFib()
}

func runFib() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	close(c)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
}

func chansize() {
	c := make(chan int, 2) // channel size 2
	c <- 1
	c <- 2
	// c <- 3 if add this line will throw error deadlock
	fmt.Println(<-c)
	fmt.Println(<-c)
	// can add 3 now
	c <- 3
	fmt.Println(<-c)

	fmt.Println("-----------")
}

func calcSum() {
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
	fmt.Println("-----------------------")
}
