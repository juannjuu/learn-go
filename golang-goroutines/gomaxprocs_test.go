package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	groups := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		groups.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCPU)

	runtime.GOMAXPROCS(20) //add threads
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads: ", totalThreads) //in default the result of threads same with CPU

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines: ", totalGoroutines)

	group.Wait()
}
