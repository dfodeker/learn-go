//go:build echo3

package echo

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[1:], " _ "))
	fmt.Println(os.Args[1:])

	//excercise 1.1 modify the echo program to also print os.Args[0],
	//the name of the command that invoked it
	fmt.Println("command name", os.Args[0])

	//Excersise 1.2 modify the echo program to print the index and value of each of its arguments one per line
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

}
