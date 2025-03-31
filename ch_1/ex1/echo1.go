package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// i := 0
	// for i < len(os.Args) {
	// 	s += sep + os.Args[i]
	// 	sep = ""
	// }

	// fmt.Println(s)
	// print(s)

	for i := 1; i < len(os.Args); i++ {

		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	print(s)
}

func print(s string) {
	fmt.Println(s)
}
