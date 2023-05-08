package main

import "fmt"

const Pi = 3.14

// var big int = 9223372036854775807 + 1
// const は overflow しない
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)
	fmt.Println(big - 1)
}
