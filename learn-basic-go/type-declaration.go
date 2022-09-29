package main

import "fmt"

func main() {
	//alias
	type noKTP string

	var noKTPJuan noKTP = "327512"
	fmt.Println(noKTPJuan)
}