// 1. create a func with the identifier foo that returns an int
// 2. create a func with the identifier bar that returns an int and a string
// 3. call both funcs
// 4. print out their results

package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Println(y, z)
}

func foo() int {
	return 42
}

func bar() (string, int){
	return "Random", 42
}
