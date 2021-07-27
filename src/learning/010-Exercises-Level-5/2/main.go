// Take the code from the previous exercise, then store the values of
// type person in a map with the key of last name. Access each value in the map.
// Print out the values, ranging over the slice.

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

	persons := map[string]person {
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, p := range persons {
		fmt.Println(k, "information:")
		fmt.Println(p.firstName, p.lastName, "favorite ice cream flavors:")

		for j, v := range p.favIceCreamFlavors {
			fmt.Printf("%d. %s\n", j + 1, v)
		}
	}
}
