package main

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	var i MyInt = 10
	println(i.Double())
}
