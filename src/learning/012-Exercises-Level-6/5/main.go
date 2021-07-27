// 1. create a type SQUARE
// 2. create a type CIRCLE
// 3. attach a method to each that calculates AREA and returns it
//		- circle area= π r 2
//		- square area = L * W
// 4. create a type SHAPE that defines an interface as anything that has the AREA method
// 5. create a func INFO which takes type shape and then prints the area
// 6. create a value of type square
// 7. create a value of type circle
// 8. use func info to print the area of square
// 9. use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}
type square struct {
	l float64
	w float64
}

func (s square) area() float64 {
	return s.l * s.w
}
func (c circle) area() float64 {
	return math.Pi * c.r * 2
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		r: 2,
	}
	s := square{
		l: 3,
		w: 2,
	}

	info(c)
	info(s)
}
