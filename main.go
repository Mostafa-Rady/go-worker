package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	totalItems := flag.Int("t", 500, "Total items to store")
	workers := flag.Int("w", 10, "Number of workers to perform in parallel")

	flag.Parse()

	fmt.Printf("Process %d items via %d workers", *totalItems, *workers)

	var processedItems int64

	startTime := time.Now()

	var wg sync.WaitGroup

	// buffer
	ch := make(chan Item, 1)

	for i := 0; i < *workers; i++ {
		go func() {
			for item := range ch {
				Do(&item)
				atomic.AddInt64(&processedItems, 1)
				wg.Done()
			}
		}()
	}

	for i := 0; i < *totalItems; i++ {
		item := Item{
			prop1: "foo",
			prop2: 1,
		}
		wg.Add(1)
		ch <- item

	}

	wg.Wait()

	fmt.Printf("Finished after: %s, processed %d items\n", time.Since(startTime), processedItems)
}

func Do(item *Item) {
	delay := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(delay)
	fmt.Print(".")
}

type Item struct {
	prop1 string
	prop2 int
}
