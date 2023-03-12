package main

import "fmt"

func main() {
	const pi = 3.14
	fmt.Println(pi)

	//error
	pi = 500
	fmt.Println(pi)
}