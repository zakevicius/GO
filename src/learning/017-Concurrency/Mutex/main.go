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
	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()
			fmt.Println("Goroutines:\t", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 50)
			runtime.Gosched()

			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

