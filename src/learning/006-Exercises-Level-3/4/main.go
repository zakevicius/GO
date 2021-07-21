// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	born := 1986

	for {
		if born > 2021 {
			break
		}
		fmt.Println(born)
		born++
	}
}
