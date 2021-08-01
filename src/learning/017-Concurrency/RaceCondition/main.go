package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	const gs = 100
	counter := 0
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			fmt.Println("Goroutines:\t", runtime.NumGoroutine())
			time.Sleep(time.Second)
			runtime.Gosched()

			v := counter
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

