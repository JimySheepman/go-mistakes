package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	for i := 0; i < 10; i++ {
		<-timer.C
		fmt.Println("timer done")
		timer.Reset(1 * time.Second)
	}
}
