package main

import "fmt"

func main() {
	m := map[string]int{
		"First": 1,
		"Second": 2,
	}

	fmt.Println("m[\"First\"]: ", m["First"])
	fmt.Println("m[\"Second\"]: ", m["Second"])
	fmt.Println("m[\"Non-existing\"]: ", m["No value"])

	fmt.Println("// Check if exists `v, ok := m[\"Non existing\"]`:")
	v, ok := m["Non existing"]
	fmt.Println("v: ", v, " ok: ", ok)

	if v, ok := m["First"]; ok {
		fmt.Println("Value exists: ", v)
	}

	fmt.Println("// Add element to map")
	fmt.Println("m[\"New\"] = 3")

	m["New"] = 3
	fmt.Println(m)

	fmt.Println("// For loop of map")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	fmt.Println("// Delete element from map `delete(map, key)`")
	delete(m, "New")
	fmt.Println(m)

	fmt.Println("// Check if exist before deleting from map")
	if _, ok := m["First"]; ok {
		delete(m, "First")
		fmt.Println(m)
	}
}
