//CIRCULAR LIST

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5)

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
