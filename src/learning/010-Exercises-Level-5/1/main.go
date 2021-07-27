// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
//		first name
//		last name
//		favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the
// elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Antoine",
		lastName: "McManocha",
		favIceCreamFlavors: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		firstName: "Shaquille",
		lastName: "Manuel",
		favIceCreamFlavors: []string{"chocolate", "pistachio"},
	}

	for _, p := range []person{p1, p2} {
		fmt.Println(p.firstName, p.lastName, "favorite ice cream flavors:")

		for j, v := range p.favIceCreamFlavors {
			fmt.Printf("%d. %s\n", j + 1, v)
		}
	}
}
