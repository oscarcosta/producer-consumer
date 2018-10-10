package main

import (
	"sync"

	log "github.com/Sirupsen/logrus"
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
	log.Info("consume: Started")
	for {
		data := c.buff.Next()
		log.Info("consume: Received ", data)
		c.wg.Done() // decrements the "done counter"
	}
}
