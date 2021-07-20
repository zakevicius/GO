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

	for x < 20 {
		if x%3 != 0 {
			x++
			continue
		}
		fmt.Println(x)
		x++
	}

	fmt.Println(`"for" with single condition`)
	x = 1

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

	fmt.Println(`### "if" initialization statement`)
	if z := 42; z == 42 {
		fmt.Println(z)
	}
	// fmt.Println(z) - undefined

	fmt.Println(`### "switch" without var after switch`)
	s := 1

	switch {
	case s == 2:
		fmt.Println("s = 2")
	case s == 1:
		fmt.Println("S = 1")
	case true:
		fmt.Println("Also true, but does not print")
	default:
		fmt.Println("This is default")
	}

	fmt.Println(`### "switch" with fallthrough`)
	switch {
	case s == 2:
		fmt.Println("s = 2")
	case s == 1:
		fmt.Println("S = 1")
		fallthrough // Opposite of JS "break", code continues
	case s == 3:
		fmt.Println("false, and also prints")
		fallthrough
	case true:
		fmt.Println("Also true, and also prints")
	case s == 4:
		fmt.Println("false, but does not print, because no fallthrough in previous case")
	}

	fmt.Println(`### "switch" with var after switch`)
	switch s {
	case 2:
		fmt.Println("s = 2")
	case 1:
		fmt.Println("S = 1")
	default:
		fmt.Println("This is default")
	}

	fmt.Println(`### "switch" with multiple cases`)
	switch s {
	case 2, 1:
		fmt.Println("s = 2 || 1")
	case 3:
		fmt.Println("S = 3")
	default:
		fmt.Println("This is default")
	}
}
