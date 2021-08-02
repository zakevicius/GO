/*
This exercise will reinforce our understanding of method sets:
	1. create a type person struct
	2. attach a method speak to type person using a pointer receiver
		- *person
	3. create a type human interface
		- to implicitly implement the interface, a human must have the speak method
	4. create func “saySomething”
		- have it take in a human as a parameter
		- have it call the speak method
	5. show the following in your code
		- you CAN pass a value of type *person into saySomething
		- you CANNOT pass a value of type person into saySomething

Here is a hint if you need some help https://play.golang.org/p/FAwcQbNtMG

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T
*/

package main

import "fmt"

type person struct {}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Speaking")
}

func main() {
	// this runs
	p1 := person{}
	saySomething(&p1)
	p1.speak()

	p2 := person{}
	// this does not run
	/*
		saySomething(p2)
	*/

	// but this do
	p2.speak() // same as (&p2).speak()
	(&p2).speak()
}

func saySomething(h human) {
	h.speak()
}
