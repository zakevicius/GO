package main

import "fmt"

func main() {
	varFoo(1, 2, 3)

	sum := varBar(1, 2, 3, 4)
	fmt.Println(`Sum:`, sum)

	avg1 := manyParams("Name1", 10, 10, 8, 8)
	fmt.Println("Average score is", avg1)
	avg2 := manyParams("Name2")
	fmt.Println("Average score is", avg2)
}

// Variadic parameter

// varFoo(x ...int) any number of ints. varFoo() also works
func varFoo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func varBar(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// Variadic parameter must be last in the list of parameters
func manyParams(name string, scores ...int) int {
	sum := 0
	avg := 0
	for _, v := range scores {
		sum += v
	}

	if len(scores) > 0 {
		avg = sum / len(scores)
	}

	fmt.Println(name)

	return avg
}

