package main

import (
	log "github.com/Sirupsen/logrus"
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
	log.Info("produce: Started")
	for i := 0; i < max; i++ {
		//time.Sleep(10 * time.Millisecond)
		log.Info("produce: Sending ", i)
		p.buff.Send(i)
	}
}
