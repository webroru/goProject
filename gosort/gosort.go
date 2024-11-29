package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers > ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	input = strings.TrimSpace(input)
	handleInput(input)
}

func handleInput(input string) {
	array := strings.Split(input, " ")
	numbers := arrayAtoi(array)
	sortedArray := sort(numbers)
	fmt.Println(sortedArray)
}

func arrayAtoi(array []string) []int {
	var numbers = []int{}
	for _, i := range array {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}
	return numbers
}

func getBatchSize(array []int) int {
	return len(array) / 4
}

func sort(array []int) []int {
	var wg sync.WaitGroup
	sortedParts := make([][]int, 4)
	batchSize := getBatchSize(array)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		start := i * batchSize
		end := start + batchSize
		if (i == 3) {
			go sortBatch(array[start:], &sortedParts[i], &wg)
		} else {
			go sortBatch(array[start:end], &sortedParts[i], &wg)
		}
	}

	wg.Wait()
	mergedArray := mergeSortedArrays(sortedParts);
	slices.Sort(mergedArray)

	return mergedArray;
}

func sortBatch(array []int, sortedPart *[]int, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine is sorting subarray: %v\n", array)
	slices.Sort(array)
	*sortedPart = array
	wg.Done()
}

func mergeSortedArrays(parts [][]int) []int {
	var result []int
	for _, v := range parts {
		result = append(result, v...)
	}

	return result
}
