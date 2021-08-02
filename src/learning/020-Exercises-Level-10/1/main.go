// fix given code using:
// 		1. func literal, aka, anonymous self-executing func
//		2. a buffered channel

/*
	func main() {
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	}
*/


package main

import (
	"fmt"
)

func main() {
	first()
	second()
}

func first() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func second() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
