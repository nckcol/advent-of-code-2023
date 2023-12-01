package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var digitMap = map[string]int{
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

func parseCalibrationNumber(source string) int {
	var digits [2]int

	for i := 0; i < len(source); i++ {
		for digit := range digitMap {
			match := true
			for j := 0; j < len(digit); j++ {
				// If we're at the end of the line or the digit doesn't match, break
				if (i+j) >= len(source) || source[i+j] != digit[j] {
					match = false
					break
				}
			}
			if match {
				if digits[0] == 0 {
					digits[0] = digitMap[digit]
				}
				digits[1] = digitMap[digit]
				break
			}
		}
	}

	return digits[0]*10 + digits[1]
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		calibrationNumber := parseCalibrationNumber(line)

		fmt.Print(calibrationNumber, " ")
		sum = sum + calibrationNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Sum:", sum)
}
