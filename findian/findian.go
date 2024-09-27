package main

import (
	"fmt"
	"strings"
)

var input string

func main() {
	fmt.Println("Enter your string")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	if strings.Index(input, "i") == 0 && strings.LastIndex(input, "n") == len(input) - 1 && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
