package main

import "fmt"

func main() {
	doit := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil
		}

		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res)
	} else {
		fmt.Println("bad result: (res is nil)")
	}
}
