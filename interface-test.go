package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("inter")
	p1 := Person{"peter", 20}

	var m Men

	m = p1
	fmt.Println(p1)
	m.sayHi()

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 100}

	t := reflect.TypeOf(m)
	tv := reflect.ValueOf(m)
	fmt.Println("t: ", t)
	fmt.Println("tv: ", tv)
	fmt.Println("tv.kind: ", tv.Kind())
	fmt.Println("tv.kind == reflect.struct: ", tv.Kind() == reflect.Struct)
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Println("x: ", x)
	fmt.Println("p: ", p)
	fmt.Println("v: ", v)

	fmt.Println("------------------")

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string { // implement person tostring method
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func (p Person) sayHi() { // must inplement this m=p1 will throw error
	fmt.Println(p.name, "say hi")
}

type Men interface {
	sayHi()
}
