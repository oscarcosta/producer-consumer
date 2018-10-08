package main

import (
	"fmt"
	"sync"
)

// Consumer definition
type Consumer struct {
	buff *Buffer
	wg   *sync.WaitGroup
}

// NewConsumer creates a Consumer
func NewConsumer(buff *Buffer, wg *sync.WaitGroup) *Consumer {
	return &Consumer{buff: buff, wg: wg}
}

// consume reads the msgs channel
func (c *Consumer) consume() {
	for {
		data := c.buff.Next()
		fmt.Printf("consume: Received %d\n", data)
		c.wg.Done() // decrements the "done counter"
	}
}
