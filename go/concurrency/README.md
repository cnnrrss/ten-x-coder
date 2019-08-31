# Concurrency in Go

**Advantages of Goroutinues over threads**
- You can run more goroutines on a typical system than you can threads.
    - Goroutines have built-in primitives to communicate safely between themselves (channels)
    - Goroutines have faster startup time
    - Goroutines are multiplexed onto a smaller number of OS threads
        - OS threads map 1:1
    - Goroutines have growable segmented stacks
    - In java you can run 1000's or 10k's threads. In Go you can run hundreds of thousands or millions of goroutines.
        - Threads consume a lot of memory (≥ 1MB).

**Closures**
- Make some concurrent calculations easier to express
- They are just local functions.
