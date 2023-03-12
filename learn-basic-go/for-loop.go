package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("counter:", counter)
	}

	slice := []string{"Juan", "Yuan", "Huan"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("slice",i, ":", slice[i])
	}

	//for range
	for index, name := range slice {
		fmt.Println("index", index, ":", name)
	}

	for _,value := range slice {
		fmt.Println(value)
	}
}