package main

import "fmt"

func main() {
	defer foo() // wait until surrounding (main()) function ends
	defer fizz() // multiple defer functions are called last-to-first
	bar()
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
