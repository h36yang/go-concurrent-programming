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
	// Instantiate channels
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// Add 2 concurrent processes to the WaitGroup
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, mx); ok {
				// Send the Book object value to the cache channel
				ch <- b
			}
			// Mark current Goroutine in the WaitGroup as done
			wg.Done()
		}(id, wg, mx, cacheCh)

		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id); ok {
				// Lock the resource for Write
				// NOTE: Only a single Writer can access the resource at one time
				mx.Lock()
				cache[id] = b
				// Unlock the resource for Write
				mx.Unlock()
				// Send the Book object value to the database channel
				ch <- b
			}
			// Mark current Goroutine in the WaitGroup as done
			wg.Done()
		}(id, wg, mx, dbCh)

		// Create one Goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) {
			// Select statement to react to whichever case triggers faster
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh // drain the msg from database channel so it doesn't trigger the second case
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

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

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond) // sleep to fake the DB call
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}
