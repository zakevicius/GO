package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// Methods for types
func (s secretAgent) speak() {
	fmt.Println("I am secret agent", s.first, s.last)
}

func (s secretAgent) shoot() {
	fmt.Println("Pew pew", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am person", p.first, p.last)
}

// Interface
type human interface {
	speak() // any type that has speak() (secretAgent && person) is also human
}

// Polymorphism
func bar(h human) {
	fmt.Println("I am BAR person", h)
}

// Assertion
func fizz(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am FIZZ person", h.(person).first)
	case secretAgent:
		fmt.Println("I am FIZZ secret Agent with license to kill: ", h.(secretAgent).first, h.(secretAgent).ltk)
	}
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last: "Moneypenny",
		},
		ltk: true,
	}

	sa1.speak()
	sa1.shoot()
	sa2.speak()
	sa2.shoot()

	p1 := person{
		first: "Dr.",
		last: "No",
	}

	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	fizz(sa1)
	fizz(sa2)
	fizz(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
