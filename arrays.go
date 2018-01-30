package main

import "fmt"

func main() {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%d - %s\n", i, fruits[i])
	}

	fmt.Println()

	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	for i, value := range fruits {
		fmt.Printf("%d - %s\n", i, value)
	}

	fmt.Println()
	// Using the keyword `continue`, skip every other fruit (even ones)

	for i, value := range fruits {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d - %s\n", i, value)
	}
}
