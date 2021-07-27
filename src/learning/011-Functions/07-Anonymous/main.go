package main

import "fmt"

func main() {
	foo()

	// anonymous function (JS - IIFE)
	func(){
		fmt.Println("Anonymous")
	}()

	func(x int){
		fmt.Println(x)
	}(42)
}

func foo() {
	fmt.Println("Foo")
}
