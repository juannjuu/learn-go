//type assertions is checking if an interface value is of a specific type

package main

import "fmt"

func random() interface{} {
	// return "Juju"
	// return 20
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is a string")
	case int:
		fmt.Println("Value", value, "is a int")
	default:
		fmt.Println("Unknown")
	}
}
