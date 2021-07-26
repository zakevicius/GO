package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type agent struct {
	person
	secret bool
	first string
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 21,
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	agent1 := agent{
		person: person{
			first: "James",
			last: "Bond",
			age: 21,
		},
		first: "First",
		secret: true,
	}
	agent2 := agent{
		person: p2,
		secret: false,
	}

	fmt.Println(agent1.first, agent1.person.first, agent1.last, agent1.age, agent1.secret)
	fmt.Println(agent2.first, agent2.last, agent2.age, agent2.secret)

	// Anonymous struct

	info := struct {
			id int
			address string
		}{
			id: 123,
			address: "Address",
		}

	fmt.Println(info)
}

