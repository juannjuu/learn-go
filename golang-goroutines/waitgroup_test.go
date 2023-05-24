package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {

	group.Add(1) // add 1 asynchronous process to group

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
	group.Done()
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Finished")
}
