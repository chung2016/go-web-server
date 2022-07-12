package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	description string
}

func main() {
	var person1 person
	fmt.Println("first time person1: ", person1)
	person1.name, person1.age = "peter", 16
	fmt.Println(person1)
	person2 := person{name: "tom", age: 18}
	fmt.Println(person2)

	student1 := student{person: person1, description: "this is student 1"}
	fmt.Println(student1)

	student1.person.age -= 1
	fmt.Println(student1)
}
