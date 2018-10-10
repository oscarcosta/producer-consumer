package main

import (
	log "github.com/Sirupsen/logrus"
)

// Producer definition
type Producer struct {
	msgs *chan int
	done *chan bool
}

// NewProducer creates a Producer
func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

// produce creates and sends the message through msgs channel
func (p *Producer) produce(max int) {
	log.Info("produce: Started")
	for i := 0; i < max; i++ {
		log.Info("produce: Sending ", i)
		*p.msgs <- i
	}
	*p.done <- true // signal when done
	log.Info("produce: Done")
}
