package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)

	fmt.Println(in, in == nil)
	in = data
	fmt.Println(in, in == nil)
	fmt.Println("in type: ", reflect.TypeOf(in))
	fmt.Println("in value:", reflect.ValueOf(in))
}

// Output:
// <nil> true
// <nil> true
// <nil> false
// in type:  *uint8
// in value: <nil>
