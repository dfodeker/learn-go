package main

import (
	"fmt"
	"os"
)

func main() {
	e := "🙈"
	fmt.Printf("oppss we have a problem %v\n", e)
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
