package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	CELL_OPERATIONAL = '.'
	CELL_DAMAGED     = '#'
	CELL_UNKNOWN     = '?'
)

type SpringRow struct {
	Row    []byte
	Chunks []int
}

func (s SpringRow) String() string {
	return fmt.Sprintf("%s %v", string(s.Row), s.Chunks)
}

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	records := make([]SpringRow, 0)
	for _, line := range lines {
		row := SpringRow{}
		parsed1 := strings.Split(line, " ")
		row.Row = []byte(parsed1[0])
		parsed2 := strings.Split(parsed1[1], ",")
		for _, chunk := range parsed2 {
			n, err := strconv.Atoi(chunk)
			if err != nil {
				log.Fatalln("Cannot parse input", err)
			}
			row.Chunks = append(row.Chunks, n)
		}
		records = append(records, row)
	}

	allCount := 0
	for _, record := range records {
		// record = unfoldSpringRow(record)
		fmt.Println(record)
		record = optimizeSpringRow(record)
		fmt.Println(record)
		count := countSpringRowArrangements(record)
		fmt.Println(count)
		allCount += count
	}

	fmt.Println("All:", allCount)
}

func unfoldSpringRow(record SpringRow) SpringRow {
	result := SpringRow{}
	for i := 0; i < 5; i++ {
		result.Row = append(result.Row, record.Row...)
		if i < 4 {
			result.Row = append(result.Row, CELL_UNKNOWN)
		}
		result.Chunks = append(result.Chunks, record.Chunks...)
	}
	return result
}

func optimizeSpringRow(record SpringRow) SpringRow {
	result := SpringRow{Chunks: record.Chunks}
	chunk := make([]byte, 0)
	for _, cell := range record.Row {
		if cell != CELL_OPERATIONAL {
			chunk = append(chunk, cell)
		} else {
			if len(chunk) > 0 {
				result.Row = append(result.Row, chunk...)
				result.Row = append(result.Row, CELL_OPERATIONAL)
				chunk = make([]byte, 0)
			}
		}
	}
	if len(chunk) > 0 {
		result.Row = append(result.Row, chunk...)
	}
	return result
}

func countSpringRowArrangements(record SpringRow) int {
	count := 0
	damagedSpringCount := 0
	unknownCount := 0
	allSpringCount := 0
	unknownIndexes := make([]int, 0)
	for index, cell := range record.Row {
		if cell == CELL_UNKNOWN {
			unknownCount++
			unknownIndexes = append(unknownIndexes, index)
		} else if cell == CELL_DAMAGED {
			damagedSpringCount++
		}
	}
	for _, chunk := range record.Chunks {
		allSpringCount += chunk
	}

	for combination := 0; combination < 1<<unknownCount; combination++ {
		row := make([]byte, len(record.Row))
		copy(row, record.Row)
		for i := 0; i < unknownCount; i++ {
			if combination&(1<<i) == 0 {
				row[unknownIndexes[i]] = CELL_DAMAGED
			} else {
				row[unknownIndexes[i]] = CELL_OPERATIONAL
			}
		}

		if isChunksEqual(getSpringRowChunks(row), record.Chunks) {
			count++
		}
	}

	return count
}

func getSpringRowChunks(row []byte) []int {
	chunks := make([]int, 0)
	chunk := 0
	for _, cell := range row {
		if cell == CELL_DAMAGED {
			chunk++
		} else {
			if chunk > 0 {
				chunks = append(chunks, chunk)
				chunk = 0
			}
		}
	}
	if chunk > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func isChunksEqual(chunks1 []int, chunks2 []int) bool {
	if len(chunks1) != len(chunks2) {
		return false
	}
	for i := 0; i < len(chunks1); i++ {
		if chunks1[i] != chunks2[i] {
			return false
		}
	}
	return true
}
