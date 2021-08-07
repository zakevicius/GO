// Package foo does foo
package foo

// Foo sums and foo
func Foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
