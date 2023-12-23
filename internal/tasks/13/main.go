package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/nckcol/advent-of-code-2023/internal/utils/fields"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	CELL_ASH  = '.'
	CELL_ROCK = '#'
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	startIndex := 0
	for startIndex < len(lines) {
		endIndex := startIndex
		for endIndex < len(lines) && lines[endIndex] != "" {
			endIndex++
		}

		field := fields.NewFromStringArray(lines[startIndex:endIndex])

		fmt.Println(field.String())

		hValues := make([]int, 0)
		vValues := make([]int, 0)

		for i := 0; i < field.Height; i++ {
			hValues = append(hValues, toNumber(field.Row(i)))
		}

		for i := 0; i < field.Width; i++ {
			vValues = append(vValues, toNumber(field.Col(i)))
		}

		hAxisIndex := findReflectionAxis(hValues, 0)
		vAxisIndex := findReflectionAxis(vValues, 0)

		hAxisIndex1 := findReflectionAxis(hValues, 1)
		vAxisIndex1 := findReflectionAxis(vValues, 1)

		partialSum := 0

		for _, h := range hAxisIndex1 {
			if h != 0 && !slices.Contains(hAxisIndex, h) {
				partialSum = 100 * h
			}
		}

		for _, v := range vAxisIndex1 {
			if v != 0 && !slices.Contains(vAxisIndex, v) {
				partialSum += v
			}
		}

		// fmt.Println(partialSum, hAxisIndex, vAxisIndex, hAxisIndex1, vAxisIndex1)

		if partialSum == 0 {
			log.Fatalf("Cannot find alternative axis for:\n%v\n", field.String())
		}
		sum += partialSum

		startIndex = endIndex + 1
	}

	fmt.Println(sum)
}

func findReflectionAxis(values []int, maxDist int) []int {
	results := make([]int, 0)

	start := make(map[int]int, 0)
	for i := 0; i < len(values)-1; i++ {
		dist := binaryDistance(values[i], values[len(values)-1])
		if dist <= maxDist && (len(values)-i)%2 == 0 {
			start[i] = 0
		}
		for j := range start {
			diff := i - j
			dist := binaryDistance(values[i], values[len(values)-1-diff])
			start[j] += dist
			if start[j] > maxDist {
				delete(start, j)
			} else if i == len(values)-2-diff {
				results = append(results, i+1)
				delete(start, j)
			}
		}

	}

	start = make(map[int]int, 0)
	for i := len(values) - 1; i > 0; i-- {
		dist := binaryDistance(values[i], values[0])
		if dist <= maxDist && (i-1)%2 == 0 {
			start[i] = 0
		}
		for j := range start {
			diff := j - i
			dist := binaryDistance(values[i], values[diff])
			start[j] += dist
			if start[j] > maxDist {
				delete(start, j)
			} else if i == diff+1 {
				results = append(results, i)
				delete(start, j)
			}
		}
	}

	return results
}

func toNumber(values []byte) int {
	n := 0
	for _, cell := range values {
		if cell == CELL_ROCK {
			n = n<<1 + 1
		} else {
			n = n << 1
		}
	}
	return n
}

func binaryDistance(a, b int) int {
	n := a ^ b
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
