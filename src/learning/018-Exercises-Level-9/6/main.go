/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
	- go run
	- go build
	- go install
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("Arch:\t", runtime.GOARCH)
	time.Sleep(time.Second * 2)
}
