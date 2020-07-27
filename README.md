# go-concurrent-programming
Practice materials for "Concurrent Programming with Go" course

### Concurrency and Parallelism
* Concurrency - have multiple tasks
* Parallelism - execute multiple tasks simultaneously

### Goroutines
A `goroutine` is a lightweight thread managed by the Go runtime.
|Threads|Goroutines|
|---:|:---|
|Have own execution stack|Have own execution stack|
|Fixed stack space (~1MB)|Variable stack space (starts @2KB)|
|Managed by OS|Managed by Go runtime|

### The Sync Package
The `sync` package (https://golang.org/pkg/sync) provides basic synchronization primitives to solve the following 2 challenges:
* Coordinating tasks - `WaitGroups`
  * A `WaitGroup` waits for a collection of goroutines to finish.
* Shared memory - `Mutexes`
  * A `mutex` is a **mut**ual **ex**clusion lock to protect shared memory.

### Channels
A `channel` is a typed conduit through which you can send and receive values. Channel allows goroutines to synchronize without explicit locks or condition variables in a loosely-coupled fashion.
* Creating a Channel
  ```Go
  ch := make(chan int)
  ```
* Creating a Buffered Channel
  ```Go
  ch := make(chan int, 5)
  ```
* Closing a Channel
  ```Go
  close(ch)
  ```

### Channel Types
Created channels are always bidirectional. The type can be restricted when the channel is passed into a function:
* Bidirectional
  ```Go
  func myFunction(ch chan int) {
      // ch here is bidirectional
      ch <- 23
      fmt.Println(<-ch)
  }
  ```
* Send-only
  ```Go
  func myFunction(ch chan<- int) {
      // ch here is send-only
      ch <- 23
  }
  ```
* Receive-only
  ```Go
  func myFunction(ch <-chan int) {
      // ch here is receive-only
      fmt.Println(<-ch)
  }
  ```

### Working with Channels in Control Flows
* If statements
  ```Go
  if msg, ok := <-ch; ok {
      // Enters here if channel is not closed, i.e. ok == true
  }
  ```
* For loops
  ```Go
  for msg := range ch {
      // Iterates until the channel is closed
  }
  ```
* Select statements
  ```Go
  ch1 := make(chan int)
  ch2 := make(chan string)
  select {
  case i := <-ch1:
      // Enters here if ch1 has a message to be received
  case ch2 <- "hello":
      // Enters here if ch2 is ready for us to send a message
  default:
      // use default case for non-blocking select
  }
  ```
