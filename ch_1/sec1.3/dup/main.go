// dup1 print the text of each line that appears more
// than once in the standard input, preceded by its count .
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

//initially I had an issue understanding how this worked and how to run it
// running the program allowd me input values but It did not mention how to actually close the input...
// its likely i missed it while i read through the chapter
//to supply input from a file - go run filename < input.txt
// using the stdin, input text and then close input with control+d
