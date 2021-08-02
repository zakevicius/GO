package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("Bye")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <- e:
			fmt.Println("Even:", v)
		case v := <- o:
			fmt.Println("Odd:", v)
		case v := <- q:
			fmt.Println("Quit:", v)
			return
		}
	}
}

