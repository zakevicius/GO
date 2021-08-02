// Get following codes running:
/*
	1.  func main() {
			cs := make(chan<- int)

			go func() {
				cs <- 42
			}()
			fmt.Println(<-cs)

			fmt.Printf("------\n")
			fmt.Printf("cs\t%T\n", cs)
		}
	2.  func main() {
			cr := make(<-chan int)

			go func() {
				cr <- 42
			}()
			fmt.Println(<-cr)

			fmt.Printf("------\n")
			fmt.Printf("cr\t%T\n", cr)
		}
*/

package main

import "fmt"

func main() {
	first()
	second()
}

func first() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func second() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
