//CONTAINER LINKED LIST IS DYNAMIC ARRAY WITH NO INDEX
//USE NEXT AND PREV CONCEPT
//HEAD AND TAIL IS NIL

package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Juan")
	data.PushBack("Timor")
	data.PushBack("Rahmadhika")
	data.PushFront("8")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Front().Prev())

	//PRINT ALL LIST
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//PRINT ALL LIST REVERSED
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
