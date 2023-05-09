// NIL CAN ONLY IMPLEMENTED TO SOME TYPE DATA
// FUNCTION, INTERFACE, MAP, SLICE, POINTER

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Juan")
	fmt.Println(person)
}
