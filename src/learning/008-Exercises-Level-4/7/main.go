// Create a slice of a slice of string ([][]string).
// Store the following data in the multi-dimensional slice:
//
//		"James", "Bond", "Shaken, not stirred"
//		"Miss", "Moneypenny", "Helloooooo, James."
//
// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	x := [][]string{[]string{"James", "Bond", "Shaken, not stirred"}, []string{"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v1 := range x {
		fmt.Println(i, ":", v1)
		for j, v2 := range v1 {
			fmt.Println("\t", j, ":", v2)
		}
	}
}
