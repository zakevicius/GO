package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age int
}

type ByName []person

func (a ByName) Len() int { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("---------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	// > GO 1.8
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

}
