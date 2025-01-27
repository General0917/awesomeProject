package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

/*
type Dog struct {
	Name string
}
*/

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriverCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	// var dog Dog = Dog{"Dog"}
	DriverCar(mike)
	DriverCar(x)
	// DriverCar(dog)
}
