package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//declare function to variable and function hasn't a name = anonymous function
	blacklist := func(name string) bool{
		return name == "joko"
	}

	registerUser("joko", blacklist)
	registerUser("juan", blacklist)

}