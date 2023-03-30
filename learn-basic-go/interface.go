package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(name HasName){
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var juan Person
	juan.Name = "Juan"

	SayHello(juan)
}