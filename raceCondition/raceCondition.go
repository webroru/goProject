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