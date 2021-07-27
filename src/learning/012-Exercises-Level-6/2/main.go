// 1. create a func with the identifier foo that
//		- takes in a variadic parameter of type int
//		- pass in a value of type []int into your func (unfurl the []int)
//		- returns the sum of all values of type int passed in
// 2. create a func with the identifier bar that
//		- takes in a parameter of type []int
//		- returns the sum of all values of type int passed in

package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)
	y := bar(xi)

	fmt.Println(x)
	fmt.Println(y)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
