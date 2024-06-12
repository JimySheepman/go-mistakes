package main

import "fmt"

type data struct {
	name string
}

func main() {
	s := []data{{"one"}}
	s[0].name = "two"
	fmt.Println(s)
}
