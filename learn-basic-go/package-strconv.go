package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	dec := strconv.FormatInt(1000000, 10) // convert to dec
	fmt.Println(dec)
	binary := strconv.FormatInt(1000000, 2) // convert to dec
	fmt.Println(binary)
	hexa := strconv.FormatInt(1000000, 16) // convert to dec
	fmt.Println(hexa)
	octal := strconv.FormatInt(1000000, 8) // convert to dec
	fmt.Println(octal)

}
