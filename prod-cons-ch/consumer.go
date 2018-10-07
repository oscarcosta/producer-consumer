package main

import "fmt"

// Consumer definition
type Consumer struct {
	msgs *chan int
}

// NewConsumer creates a Consumer
func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

// consume reads the msgs channel
func (c *Consumer) consume() {
	for {
		msg := <-*c.msgs
		fmt.Printf("consume: Received %d\n", msg)
	}
}
