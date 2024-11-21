package main

import (
	"fmt"
	"time"
)

func main() {
	sd := time.Now()
	defer func() {
		fmt.Printf("duration: %s", time.Since(sd))
	}()

	time.Sleep(time.Second * 1)
}

// Output:
// duration: 1.000942026s
