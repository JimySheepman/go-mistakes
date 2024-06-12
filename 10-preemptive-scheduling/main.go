package main

import (
	"fmt"
	"time"
)

func main() {
	// runtime.LockOSThread()
	// defer runtime.UnlockOSThread()
	// runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1 - Count:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2 - Count:", i)
			time.Sleep(150 * time.Millisecond)
			// runtime.Gosched()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main Goroutine Finished")
}
