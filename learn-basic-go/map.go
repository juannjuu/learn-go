package main

import "fmt"

func main(){
	person := map[string]string {
		"name" : "Juan",
		"gender" : "Male",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["gender"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Go-Lang"
	book["author"] = "Joko"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}