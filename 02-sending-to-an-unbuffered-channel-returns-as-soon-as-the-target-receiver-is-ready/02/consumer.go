package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Consumer
type Consumer struct {
	ingestChan chan int
	jobsChan   chan int
}

// callbackFunc is invoked each time the external lib passes an event to us.
func (c Consumer) callbackFunc(event int) {
	c.ingestChan <- event
}

// workerFunc starts a single worker function that will range on the jobsChan until that channel closes.
func (c Consumer) workerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", index)
	for eventIndex := range c.jobsChan {
		// simulate work  taking between 1-3 seconds
		fmt.Printf("Worker %d started job %d\n", index, eventIndex)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job %d\n", index, eventIndex)
	}
	fmt.Printf("Worker %d interrupted\n", index)
}

// startConsumer acts as the proxy between the ingestChan and jobsChan, with a select to support graceful shutdown.
func (c Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChan:
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}
