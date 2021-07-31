// Write and use recursive function

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(recursiveSum(a))
}

func recursiveSum(arr []int) float64 {

	if len(arr) == 0 {
		return 0
	}

	result := float64(arr[0]) + recursiveSum(arr[1:])

	return result
}
