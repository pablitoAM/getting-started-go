package main

import "fmt"

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and the value
	for key, value := range beatles {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	key := "Bob"
	if value, ok := beatles[key]; ok {
		fmt.Printf("Found beatles[%s]: %s\n", key, value)
	} else {
		fmt.Printf("Not found beatles[%s]\n", key)
	}
}
