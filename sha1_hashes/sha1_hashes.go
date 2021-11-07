package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()
	h5 := md5.New()

	h.Write([]byte(s))
	h5.Write([]byte(s))

	bs := h.Sum(nil)
	bs5 := h5.Sum(nil)

	fmt.Printf("%x\n", bs)
	fmt.Printf("%x\n", bs5)
}
