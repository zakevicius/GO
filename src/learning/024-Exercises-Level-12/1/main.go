// Get the code ready to BET on the code (add benchmarks, examples, tests). Run the following in this order:
// 	- tests
// 	- benchmarks
// 	- coverage
// 	- coverage shown in web browser
// 	- examples shown in documentation in a web browser

package main

import (
	"example.com/exercises-level-12/1/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
