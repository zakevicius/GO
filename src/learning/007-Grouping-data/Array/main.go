package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	x[2] = 42
	fmt.Println(x)

	fmt.Println("Length of x is ", len(x))
}
