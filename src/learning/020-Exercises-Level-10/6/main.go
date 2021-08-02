/*
write a program that
	- puts 100 numbers to a channel
	- pull the numbers off the channel and print them
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	receive(c)
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
