package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nckcol/advent-of-code-2023/internal/utils/tokenizer"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	lines, err := scanLines()
	if err != nil {
		log.Fatal(err)
	}

	var race Race

	for _, line := range lines {
		tokens, err := tokenizer.Tokenize(line, tokenizer.WithSeparators([]byte{':'}))
		if err != nil {
			log.Fatalln("Cannot parse input", err)
		}

		if tokens[0].Value == "Time" {
			timeString := ""
			for i := 0; i < len(tokens)-2; i += 1 {
				timeString = timeString + tokens[i+2].Value
			}
			time, err := strconv.Atoi(timeString)
			if err != nil {
				log.Fatalln("Cannot parse input", err)
			}
			race.Time = time
		} else if tokens[0].Value == "Distance" {
			distanceString := ""
			for i := 0; i < len(tokens)-2; i += 1 {
				distanceString = distanceString + tokens[i+2].Value
			}
			distance, err := strconv.Atoi(distanceString)
			if err != nil {
				log.Fatalln("Cannot parse input", err)
			}
			race.Distance = distance
		} else {
			log.Fatalln("Cannot parse input", err)
		}
	}

	fmt.Println(race)

	winning := 0
	outcomes := generateRaceOutcomes(race.Time)
	for _, outcome := range outcomes {
		if outcome > race.Distance {
			winning++
		}
	}

	fmt.Println("Result:", winning)
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

func generateRaceOutcomes(time int) []int {
	races := make([]int, 0, time)
	for acc := 0; acc < time+1; acc++ {
		races = append(races, acc*(time-acc))
	}
	return races
}
