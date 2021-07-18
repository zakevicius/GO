package main

import (
	"example.com/hello"
	"fmt"
)

func main() {
	message := hello.Hello("Random")
	fmt.Println(message)
	fmt.Println(hello.Quote())
}
