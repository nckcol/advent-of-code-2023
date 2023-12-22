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
	chunks := ""
	for index, chunk := range s.Chunks {
		if index == len(s.Chunks)-1 {
			chunks += fmt.Sprintf("%d", chunk)
		} else {
			chunks += fmt.Sprintf("%d,", chunk)
		}
	}
	return fmt.Sprintf("%s %s", string(s.Row), chunks)
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
	memo := make(map[string]int)
	for _, record := range records {
		record = unfoldSpringRow(record)
		fmt.Println(record)
		record = optimizeSpringRow(record)
		fmt.Println(record)
		count := countSpringRowArrangements(memo, record)
		fmt.Println(count)
		allCount += count
	}

	fmt.Println("All:", allCount)
	fmt.Println("Memo:", len(memo))
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

func countSpringRowArrangements(memo map[string]int, record SpringRow) int {
	// fmt.Println(record)
	if len(record.Row) == 0 {
		if len(record.Chunks) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if value, ok := memo[record.String()]; ok {
		return value
	}

	switch record.Row[0] {
	case CELL_UNKNOWN:
		row1 := make([]byte, len(record.Row))
		copy(row1, record.Row)
		row1[0] = CELL_DAMAGED
		row2 := make([]byte, len(record.Row))
		copy(row2, record.Row)
		row2[0] = CELL_OPERATIONAL
		result := countSpringRowArrangements(memo, SpringRow{Row: row1, Chunks: record.Chunks}) +
			countSpringRowArrangements(memo, SpringRow{Row: row2, Chunks: record.Chunks})
		memo[record.String()] = result
		return result
	case CELL_OPERATIONAL:
		result := countSpringRowArrangements(memo, SpringRow{Row: record.Row[1:], Chunks: record.Chunks})
		memo[record.String()] = result
		return result
	case CELL_DAMAGED:
		if len(record.Chunks) == 0 {
			memo[record.String()] = 0
			return 0
		}
		target := record.Chunks[0]
		if target > len(record.Row) {
			memo[record.String()] = 0
			return 0
		}
		for i := 0; i < target; i++ {
			if record.Row[i] == CELL_OPERATIONAL {
				memo[record.String()] = 0
				return 0
			}
		}
		if len(record.Row) > target {
			if record.Row[target] == CELL_DAMAGED {
				memo[record.String()] = 0
				return 0
			} else {
				result := countSpringRowArrangements(memo, SpringRow{Row: record.Row[target+1:], Chunks: record.Chunks[1:]})
				memo[record.String()] = result
				return result
			}
		}
		result := countSpringRowArrangements(memo, SpringRow{Chunks: record.Chunks[1:]})
		memo[record.String()] = result
		return result
	}
	log.Fatalln("Unknown cell", record.Row[0])
	return 0
}
