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
* Creating a Channel - `ch := make(chan int)`
* Creating a Buffered Channel - `ch := make(chan int, 5)`
* Closing a Channel - `close(ch)`

### Channel Types
Created channels are always bidirectional. The type can be restricted when the channel is passed into a function:
* Bidirectional - `func myFunction(ch chan int) { ... }`
* Send-only - `func myFunction(ch chan<- int) { ... }`
* Receive-only - `func myFunction(ch <-chan int) { ... }`

### Working with Channels in Control Flows
* If statements - `if msg, ok := <-ch; ok { ... }`
* For loops - `for msg := range ch { ... }`
* Select statements
  ```Go
  ch1 := make(chan int)
  ch2 := make(chan string)
  select {
  case i := <-ch1:
      ...
  case ch2 <- "hello":
      ...
  default:
      // use default case for non-blocking select
  }
  ```
