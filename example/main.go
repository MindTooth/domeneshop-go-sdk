package main

import (
	"fmt"
	"os"
	"github.com/MindTooth/domeneshop-sdk-go/domeneshop"

)

func main() {

	token := os.Getenv("DOMENESHOP_API_TOKEN")
	secret := os.Getenv("DOMENESHOP_API_SECRET")

	auth := new APIAuth(token, secret)

	fmt.Printl("Hello, %s!  How are you, %s", token, secret)
}
