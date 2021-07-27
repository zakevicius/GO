package main

import "fmt"

func main() {
	foo()
}

func foo() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	// Spreading Slice
	x := sum(xi...)
	fmt.Println(x)

	y := sum()
	fmt.Println(y)
}

func sum(x ...int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}
