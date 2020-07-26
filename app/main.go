package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// Instantiate a pointer to a WaitGroup
	wg := &sync.WaitGroup{}
	// Instantiate a pointer to a Read/Write Mutex
	mx := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// Add 2 concurrent processes to the WaitGroup
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex) {
			if b, ok := queryCache(id, mx); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			// Mark current Goroutine in the WaitGroup as done
			wg.Done()
		}(id, wg, mx)

		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex) {
			if b, ok := queryDatabase(id, mx); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
			// Mark current Goroutine in the WaitGroup as done
			wg.Done()
		}(id, wg, mx)

		time.Sleep(150 * time.Millisecond)
	}

	// Wait for all the processes in the WaitGroup to finish
	wg.Wait()
}

func queryCache(id int, mx *sync.RWMutex) (Book, bool) {
	// Lock the resource for Read
	// NOTE: Multiple Readers can access the same resource concurrently
	mx.RLock()
	b, ok := cache[id]
	// Unlock the resource for Read
	mx.RUnlock()
	return b, ok
}

func queryDatabase(id int, mx *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			// Lock the resource for Write
			// NOTE: Only a single Writer can access the resource at one time
			mx.Lock()
			cache[id] = b
			// Unlock the resource for Write
			mx.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
