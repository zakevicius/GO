package main

import "fmt"

func main() {
	n1 := factorial(4)
	fmt.Println(n1)

	n2 := factorialLoop(4)
	fmt.Println(n2)
}

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

func factorialLoop(y int) int {
	total := 1

	for ; y > 0; y-- {
		total *= y
	}

	return total
}
