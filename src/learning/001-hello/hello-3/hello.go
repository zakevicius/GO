package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello.")

	fmt.Println(n)
	foo()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("Foo.")
}

func bar() {
	fmt.Println("Bar.")
}
