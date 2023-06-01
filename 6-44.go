package main

import "fmt"

func do(i interface{}) {
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

func dos(i interface{}) {
	ss := i.(string)
	fmt.Println(ss + "!")
}

func dow(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	var i interface{} = 10
	do(i)

	var j interface{} = "hello"
	dos(j)

	dow(i)
	dow(j)
	dow(true)
}
