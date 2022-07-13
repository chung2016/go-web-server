package main

import "fmt"

type Animal struct {
	weight float64
}

type Cat struct {
	Animal
	name string
}

func (a *Animal) saySomething() {
	fmt.Println("animal say hi weight: ", a.weight)
}
func (c *Cat) saySomething() { // overwrite anonymous struct animal method
	fmt.Println("cat say hi name: ", c.name)
}

func main() {
	cat1 := Cat{Animal{100}, "haha"}
	fmt.Println(cat1)
	cat1.saySomething()
}
