package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var input string
	var names []Name

	fmt.Println("Enter file name")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}

	file, err := os.Open(input)

	if err != nil {
		fmt.Println("Error open file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		names = append(names, Name{fields[0], fields[1]})
	}

	for _, name := range names {
		fmt.Printf("First name: %s, Last name: %s\n", name.fname, name.lname)
	}
}