package main

import (
	"fmt"
)

type data struct {
}

func main() {
	a := &data{}
	b := &data{}

	if a == b {
		fmt.Printf("same address - a=%p b=%p\n", a, b)
	} else {
		fmt.Printf("different address - a=%p b=%p\n", a, b)
	}
}
