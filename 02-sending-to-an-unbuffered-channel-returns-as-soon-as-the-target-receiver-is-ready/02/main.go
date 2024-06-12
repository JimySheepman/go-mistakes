package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const workerPoolSize = 4

func main() {
	// create the consumer
	consumer := Consumer{
		ingestChan: make(chan int, 1),
		jobsChan:   make(chan int, workerPoolSize),
	}

	// Simulate external lib sending us 10 events per second
	producer := Producer{callbackFunc: consumer.callbackFunc}
	go producer.start()

	// Set up cancellation context and WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Start consumer with cancellation context passed
	go consumer.startConsumer(ctx)

	// Start workers and Add <workerPoolSize> to WaitGroup
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.workerFunc(wg, i)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	// Handle shutdown
	fmt.Println("*********************************\nShutdown signal received\n*********************************")
	cancel()  // Signal cancellation to context.Context
	wg.Wait() // Block here until are workers are done

	fmt.Println("All workers done, shutting down!")
}
