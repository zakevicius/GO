// 1. start with this slice
//		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 2. append to that slice this value
//		52
// 3. print out the slice
// 4. in ONE STATEMENT append to that slice these values
//		53
//		54
//		55
// 5. print out the slice
// 6. append to the slice this slice
//		y := []int{56, 57, 58, 59, 60}
// 7. print out the slice

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)
}
