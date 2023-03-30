/**
STRUCT IS A DATATYPE SIMILAR WITH CLASS
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

/**
STRUCT METHOD
*/
func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main(){
	var juan Customer
	juan.Name = "Juan"
	juan.Address = "Bekasi"
	juan.Age = 23

	fmt.Println(juan.Name)
	fmt.Println(juan.Address)
	fmt.Println(juan.Age)

	bebeb := Customer{
		Name: "Ananda",
		Address: "Jakarta",
		Age: 20,
	}
	fmt.Println(bebeb)

	bebeb.sayHi("juan")
}