package main

import "fmt"

func getFullName() (string, string, string) {
	return "Juan", "Timor", "Rahmadhika"
}	

func main() {
	firstName, _ , lastName := getFullName() // to ignore a return value use "_"
	fmt.Println(firstName)
	// fmt.Println(surname)
	fmt.Println(lastName)
}