// Building on the code from the previous example:
//	 1. at the package level scope, using the “var” keyword, create a VARIABLE
//		with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE
//		of your custom TYPE “x”
//	 2. in func main:
//		a. print out the value of the variable “x”
//		b. print out the type of the variable “x”
//		c. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//		d. print out the value of the variable “x”
//		e. use CONVERSION to convert the TYPE of the VALUE stored in “x” to
//		   the UNDERLYING TYPE, then:
//			 i. use the “=” operator to ASSIGN that value to “y”
//			ii. print out the value stored in “y”
// 		   iii. print out the type of “y”

package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Println("x = ", x)
	fmt.Printf("x type is %T\n", x)

	x = 420

	fmt.Println("x = ", x)

	y := int(x)

	fmt.Println("y = ", y)
	fmt.Printf("y type is %T\n", y)
}