package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 100
	fmt.Println(v.X, v.Y)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)

	v6 := new(Vertex)
	fmt.Println(v6)

	v7 := &Vertex{}
	fmt.Println(v7)

	// s := make([]int, 0)
	s := []int{}
	fmt.Println(s)

	changeVertex(v)
	fmt.Println(v)

	changeVertex2(&v)
	fmt.Println(v)
}
