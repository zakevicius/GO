package main

import "fmt"

func main() {
	var x int = 42
	y := 43
	fmt.Println(x)
	fmt.Println(y)

	sum := x + y

	fmt.Println(sum)

	x = 99

	fmt.Println(x)
}
