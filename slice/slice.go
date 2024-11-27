package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	var slice = make([]int, 3)
	var input string

	for {
		fmt.Println("Please enter a number or X for exit")
		fmt.Scan(&input)
		if input == "X" {
			fmt.Println("Good bye!")
			return
		}

		number, err := strconv.Atoi(input);
		if err != nil {
			fmt.Printf("Error: %s is not a number.\n", input)
			return
		}

		slice = append(slice, number)

		sort.Slice(slice, func(i, j int) bool {
    		return slice[i] < slice[j]
		})

		fmt.Println("Sorted slice:")
		for _, v := range slice {
		    fmt.Println(v)
		}
	}
}
