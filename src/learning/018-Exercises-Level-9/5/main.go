// Fix the race condition you created in exercise #3 by using package atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var incr int64
var wg sync.WaitGroup
const routineCount = 100

func main() {
	wg.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		go increment()
	}

	wg.Wait()
	fmt.Println("Done:", incr)
}

func increment() {
	time.Sleep(time.Millisecond * 500)
	atomic.AddInt64(&incr, 1)
	fmt.Println(atomic.LoadInt64(&incr))
	wg.Done()
}
