package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("I'm the First!")
	go fmt.Println("I'm the Second!")
	time.Sleep(1 * time.Second)
}

/* 
This Go program demonstrates how a race condition can occur when multiple goroutines execute concurrently without proper synchronization or communication. Here's a detailed breakdown:

Code Explanation
Goroutines:

go fmt.Println("I'm the First!") and go fmt.Println("I'm the Second!") are launched as separate goroutines. A goroutine is a lightweight thread of execution in Go.
These goroutines execute the fmt.Println function asynchronously.
Main Thread:

The time.Sleep(1 * time.Second) is used to pause the main thread for 1 second, ensuring the program doesn't exit immediately. This gives the goroutines time to execute.
Race Condition Explanation
A race condition happens when:

Two or more goroutines access shared resources or perform operations concurrently.
The outcome depends on the non-deterministic timing of their execution.
In this code:

The two goroutines (fmt.Println calls) compete to write to the standard output (shared resource) concurrently.
Since goroutines run asynchronously, their execution order is unpredictable. Either one could write to the output first, or their outputs might even interleave.

*/