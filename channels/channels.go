package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 1)
	fmt.Println("main...")

	go func() {
		fmt.Println("goroutine...")
		messages <- "ping"
		fmt.Println("ping done....")
	}()

	time.Sleep(time.Second)
	msg := <-messages
	fmt.Println(msg)
}
