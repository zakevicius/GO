package main

import "fmt"

func main() {
	x := 1

	fmt.Println(`"for" with single condition`)

	for x < 10 {
		fmt.Println(x)
		x *= 2
	}

	fmt.Println(`### "for" with "for clause"`)
	fmt.Println(`### for init; condition; post {}`)

	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	for i := 0; i <= 10; i++ {
		fmt.Print("| ")
		for j := 0; j < 3; j++ {
			fmt.Print(i+j, " ")
		}
	}

	fmt.Println()

	fmt.Println(`### "for" with "break"`)

	x = 1

	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x *= 2
	}

	fmt.Println(`### "for" with "continue"`)

	x = 1

	for x < 20{
		if x % 3 != 0 {
			x++
			continue
		}
		fmt.Println(x)
		x++
	}
}
