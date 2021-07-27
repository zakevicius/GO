// Create and use an anonymous struct.

package main

import "fmt"

func main() {
	anon := struct{
		value1 string
		value2 []int
		value3 [5]string
	}{
		value1: "anon string",
		value2: []int{42, 43, 44, 45, 56},
		value3: [5]string{"a", "b", "c", "d", "e"},
	}

	fmt.Println(anon)

	fmt.Println(anon.value1)
	for i, v := range anon.value2 {
		fmt.Println(i, v)
	}
	for i, v := range anon.value3 {
		fmt.Println(i, v)
	}
}
