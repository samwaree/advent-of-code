package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	var currentTotal, max int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentTotal > max {
				max = currentTotal
			}
			currentTotal = 0
		} else {
			num, _ := strconv.Atoi(line)
			currentTotal += num
		}
	}

	f.Close()

	fmt.Printf("Solution: %d", max)
}
