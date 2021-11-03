package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("goroutine...")
		msg := <-messages
		time.Sleep(time.Second)
		fmt.Println(msg)
	}()

	messages <- "ping"
	fmt.Println("main")
}
