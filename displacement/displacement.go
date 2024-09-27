package main

import (
	"fmt"
)

func main() {
	var acceleration, velocity, displacement, time float64

	fmt.Println("Enter acceleration")
	fmt.Scan(&acceleration)
	fmt.Println("Enter initial velocity")
	fmt.Scan(&velocity)
	fmt.Println("Enter initial displacement")
	fmt.Scan(&displacement)
	fmt.Println("Enter time")
	fmt.Scan(&time)

	fmt.Println(GenDisplaceFn(acceleration, velocity, displacement)(time))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*t*t/2 + v*t + s
	}
}
