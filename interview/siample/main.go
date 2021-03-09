package main

import (
	"fmt"
	"strings"
)

func alwaysFalse() bool {
	return false
}

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	fmt.Println(strings.TrimRight("CBBA", "BA"))
}
