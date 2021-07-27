package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)

	// return a function
	f := bar()
	val := f()
	fmt.Printf("Type of f is %T\n", f)
	fmt.Println(val)
	fmt.Println(bar()())
}

func foo() string {
	return "Foo"
}

func bar() func() int {
	return func() int {
		return 42
	}
}
