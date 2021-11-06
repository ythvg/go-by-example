package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var p = fmt.Println

func main() {

	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(1)
	p(string(intB))

	byt := []byte(`{"num":6.13,"str":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p(dat)

	p(dat["num"].(float64))
	fmt.Printf("%T\n", dat["num"])

	p(dat["str"].(interface{}))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	p(res)
	p(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
