package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	/*
	 * Unbuffered Channel - must have matching senders vs receivers
	 */
	ch := make(chan int)
	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		// Receive a message from the channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// Send a message to the channel
		ch <- 23
		wg.Done()
	}(ch, wg)

	/*
	 * Buffered Channel with buffer size 1
	 */
	ch2 := make(chan int, 1)
	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch2, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 23
		ch <- 42 // The second message here is not received but won't throw any panics because the channel has a buffer size to hold it
		wg.Done()
	}(ch2, wg)

	/*
	 * Send-only and Receive-only Channel Types
	 */
	ch3 := make(chan int)
	wg.Add(2)

	// This Goroutine can only receive message from the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch3, wg)

	// This Goroutine can only send message to the channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 23
		wg.Done()
	}(ch3, wg)

	wg.Wait()
}
