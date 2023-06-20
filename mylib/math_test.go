package mylib

import "testing"

// Example* を書くと godoc で Example として表示される

func Example() {
	v := Average([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(v)
}

func ExampleAverage() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
	// Output:
	// 3
}

func TestAverage(*testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
