// 1. Using a COMPOSITE LITERAL:
//		- create a SLICE of TYPE int
//		- assign 10 VALUES
// 2. Range over the slice and print the values out.
// 3. Using format printing:
//		- print out the TYPE of the slice

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Println(i, ":", v)
	}

	fmt.Printf("Type of x is `%T`", x)
}
