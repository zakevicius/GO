/*
1. Using goroutines, create an incrementer program
	- have a variable to hold the incrementer value
	- launch a bunch of goroutines
		- each goroutine should
			- read the incrementer value
			- store it in a new variable
			- yield the processor with runtime.Gosched()
			- increment the new variable
			- write the value in the new variable back to the incrementer variable
2. use waitgroups to wait for all of your goroutines to finish
3. the above will create a race condition.
4. Prove that it is a race condition by using the -race flag

If you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incr int
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
	temp := incr
	runtime.Gosched()
	temp++
	incr = temp
	fmt.Println(incr)
	wg.Done()
}
