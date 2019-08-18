# Go the language stuff


## Concurrency

Techniques, design patterns, and definitions for concurrency in Go

**Critical section**: Part of program where data that is shared between threads is acted upon

**Deadlock**: Thread 1 needs resource from thread 2, thread 2 needs resource from thread 1, neither will give the other the resource until the other one does first.
Example:
![](../assets/fig_1.1.png)


**Livelock**: when 2 threads don't advance the program
(Imagine 2 people walking towards eachother in a hallway, each move to try and get out of the others way, this goes on forever...)

## Cracking the Go Interview

Some Go solutions to problems from the famous book Cracking the Coding Interview


#### Chapter 16 Moderate
**Number swapper**
Use XOR

**Word frequencies**
Use hashmap if you think you'll do it more than once

**Intersection**
Find slope and y intercept