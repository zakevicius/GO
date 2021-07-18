package hello

import (
	"fmt"
	"rsc.io/quote"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Quote() string {
	return quote.Hello()
}
