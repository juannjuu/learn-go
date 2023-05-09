package main

import (
	"fmt"
	"os"
)

func main() {
	//OS IS RETURN COMPILED FILE LOCATION
	args := os.Args
	fmt.Println("Args:", args)

	//HOSTNAME IS RETURN HOSTNAME
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error:", err)
	}

	//GETENV IS RETURN ENVIRONMENT VARIABLE
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}