package main

import "time"

// Producer simulates an external library that invokes the
// registered callback when it has new data for us once per 100ms.
type Producer struct {
	callbackFunc func(event int)
}

func (p Producer) start() {
	eventIndex := 0
	for {
		p.callbackFunc(eventIndex)
		eventIndex++
		time.Sleep(time.Millisecond * 100)
	}
}
