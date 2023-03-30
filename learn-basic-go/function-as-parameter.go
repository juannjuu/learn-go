package main

import "fmt"

//using alias type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anas", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}