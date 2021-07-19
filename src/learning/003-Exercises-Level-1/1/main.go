// 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES
//    with the IDENTIFIERS “x” and “y” and “z”:
//		a. 42
//		b. “James Bond”
//		c. true
// 2. Now print the values stored in those variables using:
//		a. a single print statement
//		b. multiple print statements

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Printf("x = %d, type is %T\n", x, x)
	fmt.Printf("y = %s, type is %T\n", y, y)
	fmt.Printf("z = %t, type is %T\n", z, z)
}
