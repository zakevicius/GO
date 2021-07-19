// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

func main() {
	const (
		_ = iota
		a = 2021 + iota
		b = 2021 + iota
		c = 2021 + iota
		d = 2021 + iota
	)

	fmt.Println(a, b, c, d)
}
