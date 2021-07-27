package main

import "fmt"

func main() {
	foo()

	bar("Bar")

	s1 := fizz("Fizz")
	fmt.Println(s1)

	x, y := buzz("Grzegorz", "Brzęczyszczykiewicz")
	fmt.Println(x, y)

	// func expression
	f1 := func(){
		fmt.Println("Func expression")
	}
	f2 := func(x int){
		fmt.Println("Func expression parameter x", x)
	}

	f1()
	f2(42)
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("Foo")
}

func bar(s string) {
	fmt.Println(s)
}

func fizz(s string) string {
	return fmt.Sprint(s)
}

// func buzz(fn string, ln string) (string, bool) {
func buzz(fn, ln string) (string, bool) {
	s := fmt.Sprintf("%s %s,", fn, ln)
	b := false
	return s, b
}
