package main

import "fmt"

// main の前に実行される
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello, World")
}
