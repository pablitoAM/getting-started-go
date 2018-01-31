package main

import (
	"fmt"
	"sync"
)

func count(prefix string, count int, wg *sync.WaitGroup) {
	// Defer the waitgroup Done
	defer wg.Done()

	for i := 0; i < count; i++ {
		// Print out the prefix and the loop variable `i`
		fmt.Printf("[%s] %d\n", prefix, i)
	}
}

func main() {
	// Define the wait group
	var wg sync.WaitGroup

	// Increment wait group
	wg.Add(2)

	// Call count with a goroutine, first argument "first", second argument 50,
	// third argument the wait group
	go count("first", 1000000, &wg)

	// Call count with a goroutine, first argument "second", second argument 50,
	// third argument the wait group
	go count("second", 1000000, &wg)

	// Wait...
	wg.Wait()
}
