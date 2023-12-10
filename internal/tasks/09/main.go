package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
	"github.com/nckcol/advent-of-code-2023/internal/utils/list"
)

const (
	DIRECTION_LEFT  = 'L'
	DIRECTION_RIGHT = 'R'
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	historyList := make([][]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		history := make([]int, len(numbers))
		for i, number := range numbers {
			history[i], err = strconv.Atoi(number)
			if err != nil {
				log.Fatalln("Cannot parse input", err)
			}
		}
		historyList = append(historyList, history)
	}

	sumForwards := 0
	sumBackwards := 0
	for _, history := range historyList {
		sumBackwards += extrapolateHistoryBackwards(history)
		slices.Reverse(history)
		sumForwards += extrapolateHistoryForwards(history)
	}

	fmt.Println("Forwards:", sumForwards)
	fmt.Println("Backwards:", sumBackwards)
}

// a* a4 a3 a2 a1
// b* b3 b2 b1
// 0  0  0

func foldHistory(history []int) []int {
	result := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i += 1 {
		result[i] = history[i] - history[i+1]
	}
	return result
}

func foldHistoryBackwards(history []int) []int {
	result := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i += 1 {
		result[i] = history[i+1] - history[i]
	}
	return result
}

func extrapolateHistoryForwards(history []int) int {
	if list.All(history, func(n int) bool { return n == 0 }) {
		return 0
	}
	folded := foldHistory(history)
	return history[0] + extrapolateHistoryForwards(folded)
}

func extrapolateHistoryBackwards(history []int) int {
	if list.All(history, func(n int) bool { return n == 0 }) {
		return 0
	}
	folded := foldHistoryBackwards(history)
	return history[0] - extrapolateHistoryBackwards(folded)
}
