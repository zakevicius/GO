package main

import "fmt"

func main() {
	a := 42
	b := "string"
	fmt.Println(a)
	fmt.Println(b)

	aA := &a
	var bA *string = &b

	fmt.Println(aA)
	fmt.Printf("%T\n", aA)
	fmt.Println(bA)
	fmt.Printf("%T\n", bA)

	fmt.Println(*aA)
	fmt.Println(*bA)

	fmt.Println(*&a)
	fmt.Println(*&b)

	*aA = 43
	fmt.Println(a)

	*&a = 44
	fmt.Println(a)

	x := 0
	fmt.Println("x before")
	fmt.Println(&x)
	fmt.Println(x)
	bar(&x)
	fmt.Println("x after")
	fmt.Println(&x)
	fmt.Println(x)
}

func bar(y *int) {
	fmt.Println("y before")
	fmt.Println(*y)
	fmt.Println(y)
	*y = 43
	fmt.Println("y after")
	fmt.Println(*y)
	fmt.Println(y)
}
