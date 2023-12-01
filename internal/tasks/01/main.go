package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

var numberMap = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	calibrationNumbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		numberByPosition := make(map[int]int, 0)

		for source, digit := range numberMap {
			r := regexp.MustCompile(source)

			for _, match := range r.FindAllStringIndex(line, -1) {
				numberByPosition[match[0]] = digit
			}
		}

		var positions []int
		for p := range numberByPosition {
			positions = append(positions, p)
		}
		sort.Ints(positions)

		firstDigit := numberByPosition[positions[0]]
		lastDigit := numberByPosition[positions[len(positions)-1]]

		result := firstDigit*10 + lastDigit

		calibrationNumbers = append(calibrationNumbers, result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calibrationNumbers)

	sum := 0
	for _, number := range calibrationNumbers {
		sum += number
	}

	fmt.Println("Sum:", sum)
}
