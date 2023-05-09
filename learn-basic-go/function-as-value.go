//NOTES
//WE CAN STORE A FUNCTION IN A VARIABLE

package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("John"))
}
