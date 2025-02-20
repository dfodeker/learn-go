package main

import (
	"fmt"
	"time"
)

func main() {
	go moveCar("car 1")
	go moveCar("car 2")
	go moveCar("car 3")
	time.Sleep(3 * time.Second)
	fmt.Println("all cars finished")
}

func moveCar(car string) {
	fmt.Println(car, "is moving...")
	time.Sleep(2 * time.Second)
	fmt.Println(car, "has finished!")
}
