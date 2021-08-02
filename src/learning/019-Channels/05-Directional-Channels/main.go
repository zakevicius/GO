package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // send-only channel, only send values to this channel

	c <- 42
	c <- 43

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)

	d := make(<-chan int, 2) // receive-only channel, only receive values from this channel

	//d <- 42
	//d <- 43

	//fmt.Println(<-d)
	//fmt.Println(<-d)
	fmt.Println("--------")
	fmt.Printf("%T\n", d)
}

