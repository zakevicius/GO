package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(hashedPassword))

	loginPassword := `password1231`

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginPassword))
	if err != nil {
		fmt.Println("Wrong credentials")
		return
	}

	fmt.Println("Welcome")
}
