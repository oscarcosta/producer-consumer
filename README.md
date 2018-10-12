# Producer-Consumer Problem
### Two different implementations in Go to solve the Producer-Consumer Problem from The Little Book of Semaphores

**producerConsumerChannel** is using a go channel to transfer the messages from producer to consumer

**producerConsumerMutex** is has a custom Buffer to transfer the messages from producer to consumer. The Buffer is using an array a Mutex and a Cond implemented as a FIFO Queue.
