package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for r := range requests {
		<-limiter
		fmt.Println("request", r, time.Now())
	}

	burstyRequests := make(chan int, 5)
	burstyLimiter := make(chan time.Time, 3)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	for r := range burstyRequests {
		<-burstyLimiter
		fmt.Println("bursty request", r, time.Now())
	}

}
