package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nckcol/advent-of-code-2023/internal/utils/tokenizer"
)

const RED_CUBE_NUMBER = 12
const GREEN_CUBE_NUMBER = 13
const BLUE_CUBE_NUMBER = 14

const GAME_LABEL = "Game"
const RED_LABEL = "red"
const GREEN_LABEL = "green"
const BLUE_LABEL = "blue"

type Subset struct {
	Red   int
	Green int
	Blue  int
}

func parseGame(source string) (int, []Subset) {
	tokens, err := tokenizer.Tokenize(source, tokenizer.WithSeparators([]byte{':', ';', ','}))

	if err != nil {
		log.Fatalln("Cannot parse input", err)
	}

	var gameNumber int
	var subsetList []Subset
	lastSubset := Subset{}

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch {
		case token.Value == GAME_LABEL:
			{
				gameNumber, err = strconv.Atoi(tokens[i+1].Value)
				if err != nil {
					log.Fatalln("Cannot parse input", err)
				}
				i += 2
			}
		case token.Value == ";":
			{
				subsetList = append(subsetList, lastSubset)
				lastSubset = Subset{}
				i += 1
			}
		case token.Key == tokenizer.NUMBER:
			{
				count, err := strconv.Atoi(token.Value)
				if err != nil {
					log.Fatalln("Cannot parse input", err)
				}
				switch tokens[i+1].Value {
				case RED_LABEL:
					lastSubset.Red += count
				case GREEN_LABEL:
					lastSubset.Green += count
				case BLUE_LABEL:
					lastSubset.Blue += count
				default:
					log.Fatalln("Cannot parse input", err)
				}
				i += 2
			}
		default:
			i++
		}
	}

	subsetList = append(subsetList, lastSubset)

	return gameNumber, subsetList
}

func validateGame(source string) (int, bool) {
	gameNumber, subsetList := parseGame(source)

	for _, subset := range subsetList {
		if subset.Red > RED_CUBE_NUMBER || subset.Green > GREEN_CUBE_NUMBER || subset.Blue > BLUE_CUBE_NUMBER {
			return gameNumber, false
		}
	}

	return gameNumber, true
}

func calculateGamePower(source string) (int, int) {
	gameNumber, subsetList := parseGame(source)
	maxSubset := Subset{0, 0, 0}

	for _, subset := range subsetList {
		maxSubset.Red = max(maxSubset.Red, subset.Red)
		maxSubset.Green = max(maxSubset.Green, subset.Green)
		maxSubset.Blue = max(maxSubset.Blue, subset.Blue)
	}

	return gameNumber, maxSubset.Red * maxSubset.Green * maxSubset.Blue
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

	for scanner.Scan() {
		line := scanner.Text()

		gameNumber, isGameValid := validateGame(line)

		if isGameValid {
			fmt.Println(line, " - valid")
			sum1 += gameNumber
		} else {
			fmt.Println(line, " - invalid")
		}

		_, gamePower := calculateGamePower(line)

		sum2 += gamePower
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
