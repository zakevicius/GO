package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	const gs = 100
	var counter int64 = 0
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			fmt.Println("Goroutines:\t", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 50)
			runtime.Gosched()

			atomic.AddInt64(&counter, 1)
			fmt.Println("Atomic Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

