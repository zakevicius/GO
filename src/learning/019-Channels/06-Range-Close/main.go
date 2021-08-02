package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)

	// receive
	bar(c)  // no need to start goroutine because bar() waits for value from channel

	fmt.Println("Bye")
}

// send
// change to send-only channel
func foo(c chan<- int) {
	c <- 42
}

// receive
// change to receive-only channel
func bar(c <-chan int) {
	fmt.Println(<-c)
}

