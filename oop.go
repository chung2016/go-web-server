package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type ages int

type Group byte

type Person struct {
	age    ages // ages is type
	group  Group
	height float64
	weight float64
}

func (p Person) bmi() float64 {
	tmp1 := p.height / 100
	tmp2 := tmp1 * tmp1
	tmp := p.weight / tmp2
	return tmp
}

func (g Group) String() string { // seems JAVA to string
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[g]
}

func main() {

	const (
		WHITE = iota
		BLACK
		BLUE
		RED
		YELLOW
	)

	type PersonList []Person // array of persons

	fmt.Println("WHITE: ", WHITE)
	fmt.Println("BLACK: ", BLACK)
	fmt.Println("BLUE: ", BLUE)
	fmt.Println("RED: ", RED)
	fmt.Println("YELLOW: ", YELLOW)

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	c1 := Circle{10}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())

	person1 := Person{100, Group(WHITE), 170, 50}

	persons := PersonList{
		person1,
		Person{50, Group(BLACK), 150, 50},
	}

	fmt.Println("Person1 age: ", person1.age)
	fmt.Println("Person1 group: ", person1.group.String())
	fmt.Println("Person1 bmi: ", person1.bmi())
	fmt.Println(persons)
	fmt.Println("Person2 bmi: ", persons[1].bmi())
}
