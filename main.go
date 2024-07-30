package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func password_Hash(Password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 8)

	}
	return string(bytes), err
}

func check_Password(password_Hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_Hash), []byte(password))
	return err == nil
}

func main() {

	password := "Justin Dada with Selena Didi"
	hash, _ := password_Hash(password) //Not entertaining the error in this case

	fmt.Println("Actual Password : ", password)
	fmt.Println("Hashed Password : ", hash)

	matchingPasswordHash := check_Password(password, hash)
	fmt.Println("Matching Password : ", matchingPasswordHash)
}
