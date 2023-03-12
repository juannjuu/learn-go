package main

import "fmt"

//func with params
func sayHello(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

//func with return value
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else{
		return "Hello " + name
	}
}


func main() {
	sayHello("Juan", "Rahmadhika")
	fmt.Println(getHello("Juan"))
}