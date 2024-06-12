package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(0))
}
