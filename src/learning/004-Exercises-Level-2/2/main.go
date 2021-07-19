//	Using the following operators, write expressions and assign their values to variables:
//		==
//		<=
//		>=
//		!=
//		<
//		>
//	Now print each of the variables.

package main

import "fmt"

func main() {
	a := 1 == 2
	b := 3 <= 2
	c := 3 >= 3
	d := 1 != 2
	e := 3 < 2
	f := 3 > 2

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
}
