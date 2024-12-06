package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philosopherCount = 5
	maxEating = 2
	eatingRounds = 3
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id int
	leftChopstick *Chopstick
	rightChopstick *Chopstick
	eatingCount int
	eatingPermission chan bool

}

func (philosopher Philosopher) eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()

	for philosopher.eatingCount < eatingRounds {
		// Request permission to eat from the host
		host <- true
		philosopher.eatingPermission <- true

		// Acquire chopsticks
		philosopher.leftChopstick.Lock()
		philosopher.rightChopstick.Lock()

		// Start eating
		fmt.Printf("starting to eat %d\n", philosopher.id)
		time.Sleep(time.Millisecond * 100) // Simulate eating
		philosopher.eatingCount++
		fmt.Printf("finishing eating %d\n", philosopher.id)

		// Release chopsticks
		philosopher.rightChopstick.Unlock()
		philosopher.leftChopstick.Unlock()

		// Release permission to eat
		<-philosopher.eatingPermission
		<-host

		// Simulate thinking
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var wg sync.WaitGroup
	host := make(chan bool, maxEating) // Host allows at most 2 philosophers to eat
	philosophers := make([] *Philosopher, philosopherCount)
	chopsticks := make([] *Chopstick, philosopherCount)

	// Initialize chopsticks
	for i := 0; i < philosopherCount; i++ {
		chopsticks[i] = &Chopstick{}
	}

	// Initialize philosophers
	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &Philosopher{
			id: i + 1,
			leftChopstick: chopsticks[i],
			rightChopstick: chopsticks[(i + 1) % philosopherCount],
			eatingPermission: make(chan bool, 1),
		}
	}

	// Start dining
	for _, p := range philosophers {
		wg.Add(1)
		go p.eat(&wg, host)
	}

	wg.Wait()
}
