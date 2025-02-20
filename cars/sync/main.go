package main

import (
	"fmt"
	"time"
)

func main() {
	moveCar("Car 1")
	moveCar("Car 2")
	moveCar("Car 3")
}

func moveCar(name string) {
	fmt.Println(name, "is moving...")
	time.Sleep(2 * time.Second) // Simulate movement
	fmt.Println(name, "has finished!")
}
