// NOTES
// SCOPE FUNCTION CAN ACCESS VARIABLE FROM OUTSIDE

package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++ // can access variable from outside
	}

	increment()
	increment()

	fmt.Println(counter)
}
