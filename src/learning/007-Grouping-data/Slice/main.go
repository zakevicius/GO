package main

import "fmt"

func main() {
	fmt.Println("// x := []type{values} // composite literal")

	x := []int{0, 1, 2, 3, 4}
	y := []int{5, 6, 7, 8, 9}
	fmt.Println("x")
	fmt.Println(x)
	fmt.Println("Length ", len(x))

	fmt.Println("y")
	fmt.Println(y)
	fmt.Println("Length ", len(y))

	fmt.Println("// Loop a slice")
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("// Slicing a slice")
	fmt.Println("Access single element x[1]: ", x[1])
	fmt.Println("Access elements from index to index x[1:3]: ", x[1:3])
	fmt.Println("Access elements from index x[2:]: ", x[2:])
	fmt.Println("Access elements to index x[:3]: ", x[:3])

	fmt.Println("// Append to a slice")
	fmt.Println("Append single values to a slice append(x, 5, 6, 7) ", append(x, 5, 6, 7))
	fmt.Println("Append slice to a slice append(x, y...) ", append(x, y...))

	fmt.Println("// Delete from a slice")
	fmt.Println("Remove element append(x[:2], x[3:]...)", append(x[:2], x[3:]...))

	fmt.Println("// `make` a slice")
	fmt.Println("make([]int, length, capacity)")
	z := make([]int, 10, 11)
	fmt.Println("z := make([]int, 10, 11) = ", z)
	fmt.Println("len(z): ", len(z), ", cap(z): ", cap(z))
	z[0] = 0
	z[9] = 9
	fmt.Println(z)
	fmt.Println("Appending element to slice over the capacity, doubles the capacity")
	fmt.Println(z)
	fmt.Println(fmt.Println("len(z): ", len(z), ", cap(z): ", cap(z)))
	z = append(z, 10, 11)
	fmt.Println("z = append(z, 10, 11)")
	fmt.Println(z)
	fmt.Println(fmt.Println("len(z): ", len(z), ", cap(z): ", cap(z)))


	fmt.Println("Multidimensional slice multi := [][]string")
	arr1 := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(arr1)
	arr2 := []string{"Miss", "Moneypenny", "Vanilla", "Hazelnut"}
	fmt.Println(arr2)

	multi := [][]string{arr1, arr2}
	fmt.Print(multi)


}
