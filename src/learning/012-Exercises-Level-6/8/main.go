// 1. Create a func which returns a func
// 2. assign the returned func to a variable
// 3. call the returned func

package main

import "fmt"

func main() {
	bar := foo()
	res := bar(1, 2)
	fmt.Println(res)
}

func foo() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}
