// GOROUTINES RUN ASYNCHRONOUS CONCURRENTLY AND PARALLEL
// GOROUTINES MANAGED BY THREADS
package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// goroutines run asynchronously
func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Finished")

	time.Sleep(1 * time.Second) //wait to finish goroutines for 1 sec
}

func DisplayNumber(number int) {
	fmt.Println("Display Number : ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
