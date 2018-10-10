package main

import (
	log "github.com/Sirupsen/logrus"
)

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
	log.Info("consume: Started")
	for {
		msg := <-*c.msgs
		log.Info("consume: Received:", msg)
	}
}
