package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := fmt.Fprintln(os.Stdout, "Hello")
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.WriteString(os.Stdout, "Hello2")
	if err != nil {
		fmt.Println(err)
	}
}
