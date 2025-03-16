package main

import "fmt"

func main() {
	var x int = 5
	var y = x
	var z = &x

	fmt.Printf("x:%v memory:%v\n", x, &x)
	fmt.Printf("y:%v memory:%v\n", y, &y)
	fmt.Printf("z:%v\n", z)
	*z = 3
	fmt.Printf("z:%v value:%v\n", z, *z)
	fmt.Printf("x:%v memory:%v\n", x, &x)
}
