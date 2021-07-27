// Assign a func to a variable, then call that func

package main

import "fmt"

var x func() int

func main() {
	 x = func() int {
	 	return 42
	 }

	 fmt.Println(x())
}
