/*
Fix the race condition you created in the previous exercise by using a mutex
	- it makes sense to remove runtime.Gosched()
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var incr int
var wg sync.WaitGroup
var mu sync.Mutex
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
	time.Sleep(time.Millisecond * 100)
	mu.Lock()
	temp := incr

	temp++
	incr = temp

	fmt.Println(incr)
	mu.Unlock()
	wg.Done()
}
