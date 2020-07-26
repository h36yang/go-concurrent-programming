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
