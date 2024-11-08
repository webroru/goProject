package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Final counter value:", counter)
}