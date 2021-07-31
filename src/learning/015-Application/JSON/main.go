package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := person{
		First: "James",
		Last: "Bond",
		Age: 32,
	}
	p2 := person{
		First: "Miss",
		Last: "Moneypenny",
		Age: 27,
	}
	people := []person{p1, p2}

	fmt.Println(people)

	// convert to JSON
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs = []byte(s)

	var people2 []person

	err = json.Unmarshal(bs, &people2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(people2)

}
