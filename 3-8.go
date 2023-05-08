package main

import "fmt"

// var は 関数外でも宣言できる
var (
	i_3      int
	f64_3    float64
	s_3      string
	t_3, f_3 bool
)

func main() {
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t, f bool = true, false

	var (
		i_2      int
		f64_2    float64
		s_2      string
		t_2, f_2 bool
	)

	fmt.Println(i, f64, s, t, f)

	fmt.Println(i_2, f64_2, s_2, t_2, f_2)

	fmt.Printf("%T\n", f64)
}
