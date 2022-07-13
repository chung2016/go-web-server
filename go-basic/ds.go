package main

import (
	"fmt"
)

func main() {
	const (
		i      = 100
		pi     = 3.1415
		prefix = "Go_"
	)

	s, yes, number := "hello", true, 1 // string, bool, number
	s = "c" + s[2:4]                   // get s 2-4 char -> ll

	m := `hello
    %s
    %s
    world%s`
	fmt.Printf("%s\n", s)
	fmt.Println(yes)
	fmt.Println(number)
	fmt.Printf(m, "my", "awesome", "\n")

	// Array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr[1]", arr[1])
	fmt.Println("arr[2]", arr[2])

	var arr2 [10]int
	arr2[1] = 100
	fmt.Println("arr2[1]", arr2[1])
	fmt.Println("arr2[3]", arr2[3])
	fmt.Println("len(arr2)", len(arr2))

	c := [...]int{4, 5, 6}
	fmt.Println(c)

	// slice
	var ar = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	ar2 := ar[2:10] // get ar 2-10 elements
	ar2 = ar[:5]    // get ar 0-5
	ar2 = ar[5:]    // get ar 5-10
	fmt.Println(ar2)
	fmt.Println("len(ar2)", len(ar2))

	ar2 = append(ar2, "z")
	fmt.Println(ar2)

	fmt.Println("cap(ar2)", cap(ar2))

	// map : seems like js object {a: 10, b: 20}
	numbers := make(map[string]int) // declare
	numbers["1"] = 10
	numbers["100"] = 100
	fmt.Println(numbers)

	var numbers2 map[string]int
	fmt.Println("numbers2: ", numbers2)

	rating := map[string]float32{"100": 5, "aaa": 4.5, "ha": 4.5, "bb": 2} // declare

	rating["bb"] = 100

	rating["cc"] = 5555

	delete(rating, "100")
	fmt.Println(rating)

	var str string
	fmt.Println("str", str) // unassign = ""
	var integ int
	fmt.Println("integ", integ) // unassign = 0
}
