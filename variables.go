package main

import (
    "fmt"
)

func main() {
    // Declare the following variables with zero values:
    // name: command, type string
    var command string
    // name: valid, type bool
    var valid bool
    // print the value of those variables

    fmt.Printf("command: %q\n", command)
    fmt.Printf("valid: %v\n", valid)
    // declare and initialize the following variables:
    // name: foo, type: int, value: 5
    var foo int = 5
    // name bar, type: bool, value: true
    var bar bool = true
    // print the value of those variables
    fmt.Printf("foo %d\n", foo)
    fmt.Printf("bar %v\n", bar)
}
