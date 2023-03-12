package main

import "fmt"

func getFullName2()(firstName, surname, lastName string){
	firstName = "Juan"
	surname = "Perez"
	lastName = "Gonzalez"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a, b, c)
}