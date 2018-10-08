package main

import (
	"fmt"
)

// Producer definition
type Producer struct {
	buff *Buffer
}

// NewProducer creates a Producer
func NewProducer(buff *Buffer) *Producer {
	return &Producer{buff: buff}
}

// produce creates and sends the message through msgs channel
func (p *Producer) produce(max int) {
	for i := 0; i < max; i++ {
		//time.Sleep(10 * time.Millisecond)
		fmt.Printf("produce: Sending %d\n", i)
		p.buff.Send(i)
	}
}
