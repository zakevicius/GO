// Create a dog package. The dog package should have an exported func "Years" which takes human years
// and turns them into dog years (1 human year = 7 dog years)/ Document your code with comments.
// Use this code in func main()
// 	- run your program and make sure it works
// 	- run a local server with godoc and look at your documentation

package main

import (
	"example.com/exercises-level-11/dog"
	"fmt"
)

func main() {
	fmt.Println("2 human years is", dog.Years(2), "dog years.")
}
