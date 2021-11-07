package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	p := fmt.Println

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		p("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		p("subcommand 'foo'")
		p(" enable:", *fooEnable)
		p(" name:", *fooName)
		p(" tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		p("subcommand 'bar'")
		p(" level:", *barLevel)
		p(" tail:", barCmd.Args())
	default:
		p("expected 'foo' or 'bar' subcommand")
		os.Exit(1)
	}
}
