// Get the code ready to BET on (add benchmarks, examples, tests). Write a table test. Add documentation to the code. Run the following in this order:
// 	- tests
// 	- benchmarks
// 	- coverage
// 	- coverage shown in web browser
// 	- examples shown in documentation in a web browser

package main

import (
	"example.com/exercises-level-12/3/mymath"
	"fmt"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
