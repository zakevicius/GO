// Closure is when we have “enclosed” the scope of a variable in some code
// block. For this hands-on exercise, create a func which “encloses”
// the scope of a variable:

package main

import "fmt"

func main() {
	foo := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}

	a := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	fmt.Println(foo()())
}
