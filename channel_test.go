package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dhohir Pradana"
		fmt.Println("Data sent")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
