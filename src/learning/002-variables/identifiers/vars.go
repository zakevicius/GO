package main

import "fmt"

var outsideVar = 101

// custom type
type customType int

func main() {
	// short declaration operator :=
	x := 42
	y := 43
	fmt.Println(x)
	fmt.Println(y)

	sum := x + y

	fmt.Println(sum)

	x = 99

	fmt.Println(x)

	// var
	var z = 101 // auto type assignment
	var word string = "word" // defining type
	var rawString string = `"raw" string` // backticks
	var noValue int // will assign default zero_value (false, 0, "")

	fmt.Printf("z = %d\n", z)
	fmt.Printf("noValue = %d\n", noValue)
	fmt.Printf("word = %s\n", word)
	fmt.Printf("raw string = %s\n", rawString)
	foo()

	// custom type
	var random customType = 420
	fmt.Println(random)
	fmt.Printf("Type is %T\n", random)

	// Error: type mismatch.
	// var random2 int = random

	// type conversion
	var random2 int = int(random)
	fmt.Printf("Type is %T\n", random2)
}

func foo() {
	fmt.Println(outsideVar)
}
