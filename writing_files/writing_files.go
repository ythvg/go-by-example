package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := ioutil.WriteFile("./dat1", []byte("hello\ngo\n"), 0644)
	check(err)

	f, err := os.Create("./dat2")
	check(err)
	defer f.Close()

	n2, err := f.Write([]byte{115, 111, 109, 101, 10})
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
