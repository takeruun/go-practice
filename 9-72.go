package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("9-72.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	if err := ioutil.WriteFile("9-72-copy.go", content, 0644); err != nil {
		log.Fatal(err)
	}

	r := bytes.NewBuffer([]byte("abc"))
	contentN, _ := ioutil.ReadAll(r)
	fmt.Println(string(contentN))
}
