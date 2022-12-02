package main

import (
	"fmt"
	"os"
	"strings"
)

func score(opponent, me int) int {
	win := map[int]int{1: 2, 2: 3, 3: 1}
	lose := map[int]int{1: 3, 2: 1, 3: 2}
	if me == 0 {
		return lose[opponent]
	}
	if me == 3 {
		return opponent
	}
	return win[opponent]
}

func main() {
	f, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(f), "\n")

	opponentMap := map[string]int{"A": 1, "B": 2, "C": 3}
	meMap := map[string]int{"X": 0, "Y": 3, "Z": 6}
	total := 0

	for _, line := range lines {
		plays := strings.Split(line, " ")
		opponent, me := opponentMap[plays[0]], meMap[plays[1]]
		total += me + score(opponent, me)
	}

	fmt.Println(total)
}
