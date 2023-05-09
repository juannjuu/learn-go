package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()

	//WITH NO POINTER
	fmt.Println("WITH NO POINTER : ")
	fmt.Println("Host:", host)
	fmt.Println("User:", user)
	fmt.Println("Password:", password)

	//WITH POINTER
	fmt.Println("WITH POINTER : ")
	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)

}
