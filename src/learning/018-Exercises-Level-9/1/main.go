// in addition to the main goroutine, launch two additional goroutines
// 		- each additional goroutine should print something out
// 		- use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("Foo")
		time.Sleep(time.Second)
		wg.Done()
	}()

	go func() {
		fmt.Println("Bar")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Bye")
}
