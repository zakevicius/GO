package main

import (
	"example.com/021-Documentation/mymath"
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =", mymath.Sum(2, 3))
	fmt.Println("4 + 7 =", mymath.Sum(4, 7))
	fmt.Println("5 + 9 =", mymath.Sum(5, 9))
	fmt.Println("15 + 29 + 7 =", mymath.Sum(15, 29, 7))
}

