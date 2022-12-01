package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sumArray(array []int) int {
	arrSum := 0
	for _, a := range array {
		arrSum += a
	}
	return arrSum
}

func main() {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	var currentTotal int
	max := []int{0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			max = append(max, currentTotal)
			sort.Ints(max)
			max = max[1:]
			currentTotal = 0
		} else {
			num, _ := strconv.Atoi(line)
			currentTotal += num
		}
	}

	f.Close()

	fmt.Printf("Solution: %d", sumArray(max))
}
