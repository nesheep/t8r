package main

import (
	"flag"
	"strings"

	"github.com/nesheep/t8r"
)

func main() {
	flag.Parse()
	args := flag.Args()
	s := strings.Join(args, " ")
	t8r.Println(s, nil)
}
