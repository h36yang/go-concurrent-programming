# go-concurrent-programming
Practice materials for "Concurrent Programming with Go" course

### Concurrency and Parallelism
* Concurrency - have multiple tasks
* Parallelism - execute multiple tasks simultaneously

### Threads vs Goroutines
|Threads|Goroutines|
|---:|:---|
|Have own execution stack|Have own execution stack|
|Fixed stack space (~1MB)|Variable stack space (starts @2KB)|
|Managed by OS|Managed by Go runtime|

### Challenges with Concurrency
* Coordinating tasks - `WaitGroups`
  * A `WaitGroup` waits for a collection of goroutines to finish.
* Shared memory - `Mutexes`
  * A `mutex` is a **mut**ual **ex**clusion lock to protect shared memory.
* `Channels` can solve both challenges.

### Creating Channels
* Creating a Channel - `ch := make(chan int)`
* Creating a Buffered Channel - `ch := make(chan int, 5)`

### Channel Types
Created channels are always bidirectional. The type can be restricted when the channel is passed into a function:
* Bidirectional - `func myFunction(ch chan int) { ... }`
* Send-only - `func myFunction(ch chan<- int) { ... }`
* Receive-only - `func myFunction(ch <-chan int) { ... }`
