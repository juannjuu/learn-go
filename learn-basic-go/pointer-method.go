package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	juan := Man{"Juan"}
	juan.Married()

	fmt.Println(juan.Name)
}
