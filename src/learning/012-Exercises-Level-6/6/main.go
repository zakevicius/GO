// Build and use an anonymous func

package main

import "fmt"

func main() {
	func(s string) {
		fmt.Println(s)
	}("Print me")
	fmt.Println("Bye")
}
