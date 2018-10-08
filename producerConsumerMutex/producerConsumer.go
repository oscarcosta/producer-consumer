package main

import (
	"os"
	"strconv"
	"sync"
)

func main() {
	// get the maximum number of messages from arguments
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	buff := NewBuffer() // to do the communication

	// control when is all done
	var wg sync.WaitGroup
	wg.Add(max)

	// start a goroutine for Produce.produce
	go NewProducer(buff).produce(max)

	// start a goroutine for Consumer.consume
	go NewConsumer(buff, &wg).consume()

	// wait until all goroutines are done
	wg.Wait()
}
