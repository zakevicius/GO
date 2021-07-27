package main

import "fmt"

var x int

func main() {
	foo()

	fmt.Println("Increment")
	a := incrementor()
	b := incrementor()
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("b", b())
	fmt.Println("b", b())
	fmt.Println("b", b())
	fmt.Println("b", b())
}

func foo() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	bar()
	fmt.Println(x)

	var y int
	fmt.Println(y)
	y++
	fmt.Println(y)
	//bar() // breaks
	fmt.Println(y)

	{
		var z int
		fmt.Println(z)
	}
	//fmt.Println(z) // breaks
}

func bar() {
	x++
	//y++ // breaks
	fmt.Println("Foo")
}

func incrementor() func() int {
	var i int

	return func() int {
		i++
		return i
	}
}
