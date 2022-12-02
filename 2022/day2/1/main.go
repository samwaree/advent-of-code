package main

import (
	"fmt"
	"os"
	"strings"
)

func score(opponent, me int) int {
	outcome := map[int]int{1: 3, 2: 1, 3: 2}
	if outcome[me] == opponent {
		return 6
	}
	if opponent == me {
		return 3
	}
	return 0
}

func main() {
	f, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(f), "\n")

	opponentMap := map[string]int{"A": 1, "B": 2, "C": 3}
	meMap := map[string]int{"X": 1, "Y": 2, "Z": 3}
	total := 0

	for _, line := range lines {
		plays := strings.Split(line, " ")
		opponent, me := opponentMap[plays[0]], meMap[plays[1]]
		total += me + score(opponent, me)
	}

	fmt.Println(total)
}
