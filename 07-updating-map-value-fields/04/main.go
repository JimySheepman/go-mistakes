package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}
	m["x"].name = "two"
	fmt.Println(m["x"])
}
