package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Println("Worker", i, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", i, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
}
