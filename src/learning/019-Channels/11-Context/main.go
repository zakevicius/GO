package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-----------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("-----------------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("-----------------")

	ctx2 := context.Background()
	ctx2, cancel2 := context.WithCancel(ctx2)

	fmt.Println("Error check 1:\t\t", ctx2.Err())
	fmt.Println("num goroutines 1:\t", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx2.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2:\t\t", ctx2.Err())
	fmt.Println("num goroutines 2:\t", runtime.NumGoroutine())
	fmt.Println("About to cancel context")
	cancel2()
	fmt.Println("canceled context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3:\t\t", ctx2.Err())
	fmt.Println("num goroutines 3:\t", runtime.NumGoroutine())
}
