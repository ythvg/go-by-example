package main

import (
	"flag"
	"fmt"
)

func main() {
	p := fmt.Println

	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	p("word:", *wordPtr)
	p("num:", *numPtr)
	p("fork:", *forkPtr)
	p("svar:", svar)
	p("tail:", flag.Args())
}
