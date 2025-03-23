package main

import (
	"fmt"
	"time"
)

func chef(id int, orders <-chan string) {
	for order := range orders { // Receive orders
		fmt.Printf("⏲️ Chef %d is making %s\n", id, order)
		time.Sleep(time.Second) // Simulate cooking time
		fmt.Printf("🍕 Chef %d finished %s\n", id, order)
	}
}

func main() {
	orders := make(chan string, 3) // Create an order queue

	// Start 2 chefs (goroutines)
	for i := 1; i <= 2; i++ {
		go chef(i, orders)
	}

	// Send orders
	menu := []string{"Pepperoni", "Veggie", "BBQ Chicken", "Hawaiian", "Philly Steak"}
	for _, order := range menu {
		orders <- order
	}

	close(orders)               // Close the channel when done
	time.Sleep(3 * time.Second) // Give chefs time to finish
}
