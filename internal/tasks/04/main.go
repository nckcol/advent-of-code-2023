package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nckcol/advent-of-code-2023/internal/utils/tokenizer"
)

type Card struct {
	id               int
	winningNumbers   []int
	availableNumbers []int
}

func parseCard(source string) Card {
	tokens, err := tokenizer.Tokenize(source, tokenizer.WithSeparators([]byte{':', '|'}))

	if err != nil {
		log.Fatalln("Cannot parse input", err)
	}

	card := Card{}
	isSecondHalf := false

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch {
		case token.Value == "Card":
			{
				card.id, err = strconv.Atoi(tokens[i+1].Value)
				if err != nil {
					log.Fatalln("Cannot parse input", err)
				}
				i += 2
			}
		case token.Value == "|":
			{
				isSecondHalf = true
				i += 1
			}
		case token.Key == tokenizer.NUMBER:
			{
				count, err := strconv.Atoi(token.Value)
				if err != nil {
					log.Fatalln("Cannot parse input", err)
				}
				if isSecondHalf {
					card.availableNumbers = append(card.availableNumbers, count)
				} else {
					card.winningNumbers = append(card.winningNumbers, count)
				}
				i += 1
			}
		default:
			i++
		}
	}

	return card
}

func calculateMatches(card Card) int {
	matchCount := 0
	winningMap := make(map[int]bool, len(card.winningNumbers))

	for _, number := range card.winningNumbers {
		winningMap[number] = true
	}

	for _, number := range card.availableNumbers {
		if _, ok := winningMap[number]; ok {
			matchCount += 1
		}
	}

	return matchCount
}

func calculateCardValue(card Card) int {
	matchCount := calculateMatches(card)

	if matchCount == 0 {
		return 0
	}

	return 1 << (matchCount - 1)
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	sum1 := 0
	sum2 := 0
	wonCards := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		card := parseCard(line)
		wonCards[card.id] += 1

		value := calculateCardValue(card)
		sum1 += value

		matchCount := calculateMatches(card)
		for i := 0; i < matchCount; i++ {
			wonCards[card.id+i+1] += wonCards[card.id]
		}
		sum2 += wonCards[card.id]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
