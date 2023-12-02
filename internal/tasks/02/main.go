package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const RED_CUBE_NUMBER = 12
const GREEN_CUBE_NUMBER = 13
const BLUE_CUBE_NUMBER = 14

const RED_LABEL = "red"
const GREEN_LABEL = "green"
const BLUE_LABEL = "blue"

type Subset struct {
	Red   int
	Green int
	Blue  int
}

func parseGame(source string) (int, []Subset) {
	subsetList := make([]Subset, 0)
	split1 := strings.Split(source, ": ")
	split2 := strings.Split(split1[0], " ")
	gameNumber, err := strconv.Atoi(split2[1])

	if err != nil {
		log.Fatalln("Cannot parse input", err)
	}

	split3 := strings.Split(split1[1], "; ")

	for _, subsetSource := range split3 {
		split4 := strings.Split(subsetSource, ", ")
		subset := &Subset{}

		for _, cubeSource := range split4 {
			split5 := strings.Split(cubeSource, " ")
			count, err := strconv.Atoi(split5[0])
			if err != nil {
				log.Fatalln("Cannot parse input", err)
			}

			switch split5[1] {
			case RED_LABEL:
				subset.Red += count
			case GREEN_LABEL:
				subset.Green += count
			case BLUE_LABEL:
				subset.Blue += count
			default:
				log.Fatalln("Cannot parse input", err)
			}
		}

		subsetList = append(subsetList, *subset)
	}

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
