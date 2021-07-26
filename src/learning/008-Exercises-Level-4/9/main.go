// Using the code from the previous example, add a record to your map.
// Now print the map out using the “range” loop

package main

import "fmt"

func main() {
	x := map[string][]string {
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["random_dude"] = []string{"something1", "Something2"}

	for k, v1 := range x {
		fmt.Println(k, ":")
		for i, v2 := range v1 {
			fmt.Println(i, ":", v2)
		}
	}
}
