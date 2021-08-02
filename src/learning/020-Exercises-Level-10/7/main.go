/*
write a program that
	- launches 10 goroutines
	- each goroutine adds 10 numbers to a channel
	- pull the numbers off the channel and print them
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go channelPower(c, i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func channelPower(c chan<- int, n int) {
	wg.Add(1)
	for i := 1; i <= 10; i++ {
		c <- n * i
	}
	wg.Done()
}

