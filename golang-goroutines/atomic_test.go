//WE CAN USE ATOMIC TO MANIPULATE PRIMITIVE DATA INSTEAD OF USING LOCK MUTEX

package golanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1) //use atomic to manipulate primitive data (int64)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", x)
}
