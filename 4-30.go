package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("4-30.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatal(err)
	}
}
