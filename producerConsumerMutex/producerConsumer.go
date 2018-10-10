package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// profile flags
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	debug := flag.Bool("debug", false, "Display debug logs")

	// get the maximum number of messages from flags
	max := flag.Int("n", 0, "defines the number of messages")

	flag.Parse()

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// log level
	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	// CPU Profile
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	buff := NewBuffer() // to do the communication

	// control when is all done
	var wg sync.WaitGroup
	wg.Add(*max)

	// start a goroutine for Produce.produce
	go NewProducer(buff).produce(*max)

	// start a goroutine for Consumer.consume
	go NewConsumer(buff, &wg).consume()

	// wait until all goroutines are done
	wg.Wait()

	// Memory Profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
