package main

import "fmt"

// the classic fizzbuzz game
// prints the numbers 1-100  inclusive each on their own line, but substitues multipes of 3 for the text fizz and nultiples of 5 for buzz
// for multiples of 3 and 5 print fizzbuzz
func fizzbuzz(num int) {
	for i := 0; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz(100)
}
