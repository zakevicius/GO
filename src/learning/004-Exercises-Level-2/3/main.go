// Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

func main() {
	const (
		typed string = "is TYPED"
		untyped = "is UNTYPED"
	)

	fmt.Println("typed", typed)
	fmt.Println("untyped", untyped)
}
