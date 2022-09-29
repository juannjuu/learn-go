package main

import "fmt"

func main() {
	var names[2] string

	names[0] = "Juan"
	names[1] = "Rahmadhika"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var numbers = [3] int{
		1, 
		2, 
		3,
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	//redeclare var Numbers
	var newNumbers[10] int

	fmt.Println(len(newNumbers))
}