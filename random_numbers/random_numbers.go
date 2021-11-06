package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Print

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	p(r3.Intn(100), ",")
	p(r3.Intn(100))
	fmt.Println()

}
