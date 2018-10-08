package main

import "sync"

// Buffer deffinition
// implemented as a FIFO queue, based on https://github.com/golang/exp/blob/master/shiny/driver/internal/event/event.go
type Buffer struct {
	mu   sync.Mutex
	cond sync.Cond
	data []int
}

// NewBuffer creates a new buffer
func NewBuffer() *Buffer {
	buff := &Buffer{data: make([]int, 0)}
	buff.cond = *sync.NewCond(&buff.mu)
	return buff
}

// Send adds an value to the end of the queue
func (b *Buffer) Send(i int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.data = append(b.data, i)

	b.cond.Signal()
}

// Next returns the next value in the queue
func (b *Buffer) Next() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for {
		if len(b.data) > 0 {
			d := b.data[0]
			b.data = b.data[1:]
			return d
		}
		b.cond.Wait()
	}
}
