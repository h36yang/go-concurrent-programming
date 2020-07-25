# go-concurrent-programming
Practice materials for "Concurrent Programming with Go" course

### Concurrency and Parallelism
* Concurrency - have multiple tasks
* Parallelism - execute multiple tasks simultaneously

### Threads vs Goroutines
|Threads|Goroutines|
|---:|---|
|Have own execution stack|Have own execution stack|
|Fixed stack space (~1MB)|Variable stack space (starts @2KB)|
|Managed by OS|Managed by Go runtime|
