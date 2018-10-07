package main

import (
	"os"
	"strconv"
)

func main() {
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	var msgs = make(chan int)  // channel to send messages
	var done = make(chan bool) // channel to control when production is done

	// start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(max)

	// start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()

	// finish the program when the production is done
	<-done
}
