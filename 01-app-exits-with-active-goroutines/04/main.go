package main

import (
	"fmt"
	"time"
)

func main() {
	workerCount := 2
	done := make(chan bool, workerCount)

	for i := 0; i < workerCount; i++ {
		go doit(i, done)
	}

	for i := 0; i < workerCount; i++ {
		<-done
	}

	fmt.Println("all done!")
}

func doit(workerId int, done chan bool) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(2 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
	done <- true
}
