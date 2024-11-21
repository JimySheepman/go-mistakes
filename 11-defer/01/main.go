package main

import (
	"fmt"
	"time"
)

func main() {
	sd := time.Now()
	defer fmt.Printf("duration: %s", time.Since(sd))

	time.Sleep(time.Second * 1)
}

// Output:
// duration: 113ns
