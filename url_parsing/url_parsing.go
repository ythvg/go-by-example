package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)

	p(u.User)
	p(u.User.Username())
	p(u.User.Password())

	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host, port)

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}
