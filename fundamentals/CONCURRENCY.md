# Concurrency

#### Processes & Threads
**Process** - two separate programs
    - Can only shary data over protocol connection

**Threads** - two _channels_ of work in a program that can share data

### Channels golang
- Primary method of communication between goroutines

#### Mutexes / Locks
- Mutex locks access between shared threads to a shared resource.
- Ensure exclusive acess to a resource
- Careful can slow program

#### Semaphores & Monitors
- Semaphore: int to control the number of resources available. (think 4 cash registers).
- Monitor: abstract datatype that allows only one process at a time

#### Deadlock & Livelock
- Deadlock: multiple threads of a program waiting for the other to finish. Blocked for ever.
- Livelock: two resources continually try to get out of each others way but without moving the process forward. (Two people coming at each other in hallway)

#### Resources of a process
- CPU/Memory/Storage/etc..

#### Scheduling

**Worker Pools:**

**Wait Groups:**

**Atomic Counters:**
Update shared resource concurrently `AddInt(&num, 1)`
To extract a value being concurrently updated we use `LoadInt(&num)`

Dining philosophers