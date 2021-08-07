package main

import (
	"example.com/023-Testing-and-benchmarking/Coverage/saying"
	"fmt"
)

func main() {
	fmt.Println(saying.Greet("Blah"))
}


// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
