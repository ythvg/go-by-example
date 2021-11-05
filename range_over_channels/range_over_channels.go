package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	limiter := time.Tick(200 * time.Millisecond)

	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		<-limiter
		fmt.Println(elem)
	}
}
