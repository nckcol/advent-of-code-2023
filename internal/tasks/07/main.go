package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cardValue = map[byte]int{
	'A': 0xf,
	'K': 0xe,
	'Q': 0xd,
	'J': 0xc,
	'T': 0xb,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardValueWithJoker = map[byte]int{
	'A': 0xf,
	'K': 0xe,
	'Q': 0xd,
	'T': 0xb,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Bid struct {
	Hand      []byte
	HandValue int
	Bid       int
}

func main() {
	ensurePipeInput()
	lines, err := scanLines()
	if err != nil {
		log.Fatal(err)
	}

	var set []*Bid

	for _, line := range lines {
		input := strings.Split(line, " ")
		hand := []byte(input[0])
		bid, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatalln("Cannot parse input", err)
		}
		set = append(set, &Bid{Hand: hand, Bid: bid})
	}

	for _, bid := range set {
		bid.HandValue = calculateHandValue(bid.Hand)
	}

	slices.SortFunc(set, func(a, b *Bid) int { return a.HandValue - b.HandValue })

	result := 0
	for rank, bid := range set {
		score := bid.Bid * (rank + 1)
		// fmt.Printf("%d\t%s\t%x\t%d\n", rank+1, string(bid.Hand), bid.HandValue, score)
		result += score
	}

	fmt.Println("Result 1:", result)

	for _, bid := range set {
		bid.HandValue = calculateHandValueWithJoker(bid.Hand)
	}

	slices.SortFunc(set, func(a, b *Bid) int { return a.HandValue - b.HandValue })

	result = 0
	for rank, bid := range set {
		score := bid.Bid * (rank + 1)
		// fmt.Printf("%d\t%s\t%x\t%d\n", rank+1, string(bid.Hand), bid.HandValue, score)
		result += score
	}

	fmt.Println("Result 2:", result)
}

func ensurePipeInput() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}
}

func scanLines() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func calculateHandValue(hand []byte) int {
	value := 0
	kindCount := make(map[byte]int, 0)
	countStats := make(map[int]int, 0)
	counts := make([]int, 5)
	for _, card := range hand {
		kindCount[card] += 1
	}
	for _, count := range kindCount {
		countStats[count] += 1
		counts = append(counts, count)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	value = 0x10*(counts[0]) + counts[1]

	for _, card := range hand {
		value = value*0x10 + cardValue[card]
	}

	return value
}

func calculateHandValueWithJoker(hand []byte) int {
	value := 0
	kindCount := make(map[byte]int, 0)
	countStats := make(map[int]int, 0)
	counts := make([]int, 5)
	for _, card := range hand {
		kindCount[card] += 1
	}
	for card, count := range kindCount {
		if card == 'J' {
			continue
		}
		countStats[count] += 1
		counts = append(counts, count)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	value = 0x10*(counts[0]+kindCount['J']) + counts[1]

	for _, card := range hand {
		value = value*0x10 + cardValueWithJoker[card]
	}

	return value
}
