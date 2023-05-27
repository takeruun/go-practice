package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println(*p)

	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	// point を返す
	p3 := new(int)
	fmt.Printf("%T\n", p3)

	// point を返す
	st := new(struct{})
	fmt.Printf("%T\n", st)

}
