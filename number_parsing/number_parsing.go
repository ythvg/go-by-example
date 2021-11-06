package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := fmt.Println

	p(strconv.ParseFloat("1.234", 64))

	p(strconv.ParseInt("0x1c8", 0, 64))

	p(strconv.ParseUint("789", 0, 64))

	p(strconv.Atoi("145"))
}
