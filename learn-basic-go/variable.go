package main

import "fmt"

func main() {
	//data type of variable without a data have to be declarated 
	var name string

	name = "juannjuu"
	fmt.Println(name)

	name ="anandahrns"
	fmt.Println(name)

	var friendName = "Adam"
	fmt.Println(friendName)

	// ":=" is replace for var in the first declaration 
	country := "Indonesia"
	fmt.Println(country)

	//multiple variable
	var (
		firstName = "Juan"
		lastName = "Rahmadhika"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}