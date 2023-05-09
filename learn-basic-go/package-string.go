package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Juan Rahmadhika", "Juan")) //return true
	fmt.Println(strings.Contains("Juan Rahmadhika", "Jaja")) //return false

	fmt.Println(strings.Split("Juan Rahmadhika", " ")) //return slice [Juan Rahmadhika]

	fmt.Println(strings.ToLower("Juan Rahmadhika")) //return juan rahmadhika
	fmt.Println(strings.ToUpper("Juan Rahmadhika")) //return JUAN RAHMADHIKA
	fmt.Println(strings.ToTitle("Juan Rahmadhika")) //return JUAN RAHMADHIKA

	fmt.Println(strings.Trim("     Juan Rahmadhika      ", " "))      //cut start and end of strings, return Juan Rahmadhika
	fmt.Println(strings.ReplaceAll("Juan Juan Juan", "Juan", "Juju")) //return Juju Juju Juju

}
