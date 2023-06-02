package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// String() は fmt.Stringer インターフェースを実装する
// Mike (25 years) のように出力される
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	mike := Person{"Mike", 25}
	fmt.Println(mike)
}
