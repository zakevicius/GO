// Create a variable of type string using a raw string literal. Print it.

package main

import "fmt"

func main() {
	var a string = `Raw 
		"useless" string`

	fmt.Println(a)
}
