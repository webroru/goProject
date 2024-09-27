package main

import "fmt"

var number float32

func main() {
	fmt.Println("Enter float number")
	fmt.Scan(&number)
	fmt.Printf("Truncated value is %d\n", int32(number))
}
