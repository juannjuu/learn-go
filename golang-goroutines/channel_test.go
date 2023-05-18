// CHANNEL PASS DATA BY REFERENCE SO DIDNT NEED TO USE POINTER
// ALWAYS CLOSE CHANNEL IN THE END OR CAN CAUSE ERROR/FOREVER LOOP
package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Juan Rahmadhika"
		fmt.Println("Finished send data to channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
	close(channel)
}

func GiveResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Juan Rahmadhika"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveResponse(channel)

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
	close(channel)
}

// only use to send data to channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Send Data To Channel"
}

// only use to receive data from channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}

// buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	go func() {
		channel <- "Juan"
		channel <- "Rahmadhika"
		channel <- "HAHA"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Finished Buffer")
	time.Sleep(2 * time.Second)

	close(channel)
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Iterate : " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveResponse(channel1)
	go GiveResponse(channel2)

	counter := 0 //use counter to avoid single select case
	for {
		select {
		case data := <-channel1:
			fmt.Println("Receive data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Receive data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting for data")
		}

		if counter == 2 {
			break
		}
	}
	close(channel1)
	close(channel2)
}
