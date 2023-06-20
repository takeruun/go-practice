package mylib

// Average returns the average of a series of numbers
func Average(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum / len(s)
}
