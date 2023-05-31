package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(h Human) {
	if h.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	// Person は Say() というメソッドを持ってないといけない
	var mike Human = &Person{"Mike"}
	DriveCar(mike)

	var x Human = &Person{"X"}
	DriveCar(x)

	// var dog Dog = Dog{"dog"}
	// Dog does not implement Human (missing Say method)
	// DriveCar(dog)
}
