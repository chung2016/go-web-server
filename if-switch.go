package main

import "fmt"

func main() {
	var x int = 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	y := 3
	y = 100
	if y == 3 {
		fmt.Println("The y is equal to 3")
	} else if y < 3 {
		fmt.Println("The y is less than 3")
	} else {
		fmt.Println("The y is greater than 3")
	}
	myFuncGoto()
	loop()
	loop2()
	loop3()
	mySwitch()
	mySwitch2()
}

func myFuncGoto() {
	i := 0
Here:
	println(i)
	i++
	if i < 10 {
		goto Here
	} else {
		fmt.Println("end myFnc")
	}
}

func loop() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("loop sum: ", sum)

}

func loop2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	print("loop2 sum: ", sum)
}

func loop3() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Println("loop3 index: ", index)
	}
}

func mySwitch() {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

func mySwitch2() {
	i := 6
	switch i {
	case 4:
		fmt.Println("The i was <= 4")
		fallthrough
	case 5:
		fmt.Println("The i was <= 5")
		fallthrough
	case 6:
		fmt.Println("The i was <= 6")
		fallthrough
	case 7:
		fmt.Println("The i was <= 7")
		fallthrough
	case 8:
		fmt.Println("The i was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
