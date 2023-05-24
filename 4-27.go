package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "wwwwwwwww"
}

func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!!")
	case t.Hour() < 17:
		fmt.Println("Afternoon!!")
	case t.Hour() < 20:
		fmt.Println("Evening!!")
	default:
		fmt.Println("Night!!")
	}
}
