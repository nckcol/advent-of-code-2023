package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/nckcol/advent-of-code-2023/internal/utils/fields"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	CELL_ROUND_ROCK = 'O'
	CELL_CUBE_ROCK  = '#'
	CELL_EMPTY      = '.'
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	field := fields.NewFromStringArray(lines)

	rocks := field.FindAllPositions(func(c byte) bool { return c == CELL_ROUND_ROCK })

	for i := 0; i < 1000; i++ {
		prev := make([]fields.FieldPosition, len(rocks))
		copy(prev, rocks)
		rocks = move(field, rocks, fields.DIRECTION_NORTH)
		rocks = move(field, rocks, fields.DIRECTION_WEST)
		rocks = move(field, rocks, fields.DIRECTION_SOUTH)
		rocks = move(field, rocks, fields.DIRECTION_EAST)
		if equals(prev, rocks) {
			fmt.Println("Loop detected:", i)
			break
		}
	}

	fmt.Println(field.String())
	fmt.Println(calculateLoad(field, rocks))
}

func move(field *fields.Field, rocks []fields.FieldPosition, direction fields.Direction) []fields.FieldPosition {
	for _, rock := range rocks {
		field.Set(rock, CELL_EMPTY)
	}

	slices.SortFunc(rocks, func(a, b fields.FieldPosition) int {
		switch direction {
		case fields.DIRECTION_NORTH:
			return a.Y - b.Y
		case fields.DIRECTION_WEST:
			return a.X - b.X
		case fields.DIRECTION_SOUTH:
			return b.Y - a.Y
		case fields.DIRECTION_EAST:
			return b.X - a.X
		}
		log.Fatalln("Unknown direction", direction)
		return 0
	})

	resultedRocks := make([]fields.FieldPosition, 0, len(rocks))
	for _, rock := range rocks {
		next := rock
		for field.Available(next.Move(direction)) && field.Get(next.Move(direction)) == CELL_EMPTY {
			next = next.Move(direction)
		}
		resultedRocks = append(resultedRocks, next)
		field.Set(next, CELL_ROUND_ROCK)
	}

	return resultedRocks
}

func calculateLoad(field *fields.Field, rocks []fields.FieldPosition) int {
	load := 0
	for _, rock := range rocks {
		load += field.Height - rock.Y
	}
	return load
}

func equals(a, b []fields.FieldPosition) bool {
	if len(a) != len(b) {
		return false
	}
	bMap := make(map[fields.FieldPosition]bool)
	for _, p := range a {
		if !bMap[p] {
			return false
		}
	}
	return true
}
