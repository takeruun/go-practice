package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6}

	small := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] < small {
			small = l[i]
		}
	}

	fmt.Println(small)

	var nim int
	for i, num := range l {
		if i == 0 || num < nim {
			nim = num
		}
	}
	fmt.Println(nim)

	m := map[string]int{
		"apple":  100,
		"banana": 200,
		"lemon":  300,
		"orange": 400,
		"papaya": 30,
	}

	total := 0
	for _, v := range m {
		total += v
	}

	fmt.Println(total)

}
