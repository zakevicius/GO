package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(xi...)
	fmt.Println("All", s)

	es := even(sum, xi...)
	fmt.Println("Even", es)

	os := odd(sum, xi...)
	fmt.Println("Odd", os)
}
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var ii []int

	for _, v := range vi {
		if v%2 == 0 {
			ii = append(ii, v)
		}
	}

	return f(ii...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var ii []int

	for _, v := range vi {
		if v%2 != 0 {
			ii = append(ii, v)
		}
	}

	return f(ii...)
}
