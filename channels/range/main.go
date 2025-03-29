package main

import "fmt"

func concurrentFib(n int) []int {
	arr := make([]int, 0)
	ch := make(chan int)
	go func() {
		fibonacci(n, ch)

	}()

	for i := range ch {
		fmt.Println(i)

		arr = append(arr, i)
	}
	return arr
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
