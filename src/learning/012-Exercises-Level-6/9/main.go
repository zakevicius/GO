// pass a func into a func as an argument

package main

import "fmt"

func main() {
	f := func(times int) float64 {
		return 3.14 * float64(times)
	}

	area := area(f, 7)
	fmt.Println(area)
}

func area(f func(times int) float64, r float64) float64 {
	return f(2) * r
}
