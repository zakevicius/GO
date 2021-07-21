// Create a program that uses a switch statement with the switch expression
// specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	var favSport string

	switch favSport {
	case "baseball":
		fmt.Println("false")
	case "hockey":
		fmt.Println("true")
	default:
		fmt.Println("curling")
	}
}
