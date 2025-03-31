// modify dup2 to print the naames of all wiles in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Printf(" file name %v", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}

	}

	for line, n := range counts {
		if n > 1 {
			for _, arg := range files {
				fmt.Printf("%d\t%s\t%v \n", n, line, arg)
			}

		}
	}

}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// note ignoring potential errors from input. Err
	//the book mentions error handeling in section 5

}

/**
The function os.open returns two valeus , the first is an open file of type os.File that is used
in subsequent reads by the scanner
the second result  is a value of the built in error type , if the err wquals the special built in value nin the file was
open succesfullly , the file is read and when the end of the input is reached, Close closes the file and releases any resources.
on the other hand iff the err is not nil
something went wrong. in that case the error valuse describes the problem.
*/
