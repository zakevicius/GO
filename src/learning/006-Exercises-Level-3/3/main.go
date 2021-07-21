// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	born := 1986

	for born <= 2021 {
		fmt.Println(born)
		born++
	}
}
