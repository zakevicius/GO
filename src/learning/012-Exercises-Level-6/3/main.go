// Use the “defer” keyword to show that a deferred func runs after the func
// containing it exits.

package main

import "fmt"

func main() {
	foo()
	defer bar()
	fizz()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

func fizz() {
	fmt.Println("Fizz")
}
