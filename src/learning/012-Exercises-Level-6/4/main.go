// 1. Create a user defined struct with
//		- the identifier “person”
//		- the fields:
//			first
//			last
//			age
// 2. attach a method to type person with
//		- the identifier “speak”
//		- the method should have the person say their name and age
// 3. create a value of type person
// 4. call the method from the value of type person

package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p1 := person {
		first: "Name",
		last: "Surname",
		age: 25,
	}

	p1.speak()
}
