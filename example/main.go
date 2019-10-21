package main

import (
	"fmt"
	"os"
)

func main() {

	tokenp, err := os.LookupEnv("DOMENESHOP_API_TOKEN")

	if !err {
		fmt.Println("Testing")
		panic(err)
	}

	secretp, err := os.LookupEnv("DOMENESHOP_API_SECRET")

	if !err {
		fmt.Println("Heo")
		panic(err)
	}

	// var auth *APIAuth
	// auth.token = token
	// auth.secret = secret

	// auth := domeneshop.APIAuth{
	// 	Token:  tokenp,
	// 	Secret: secretp,
	// }

	fmt.Printf("Hello, %s!  How are you, %s", tokenp, secretp)
}
