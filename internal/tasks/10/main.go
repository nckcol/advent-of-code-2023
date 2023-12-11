package main

import (
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/nckcol/advent-of-code-2023/internal/utils/fields"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	PIPE_NORTH_SOUTH = '|'
	PIPE_EAST_WEST   = '-'
	PIPE_NORTH_EAST  = 'L'
	PIPE_NORTH_WEST  = 'J'
	PIPE_SOUTH_WEST  = '7'
	PIPE_SOUTH_EAST  = 'F'
	PIPE_START       = 'S'
	GROUND           = '.'
)

var (
	PIPES_SOUTH = []byte{PIPE_NORTH_SOUTH, PIPE_SOUTH_EAST, PIPE_SOUTH_WEST}
	PIPES_NORTH = []byte{PIPE_NORTH_SOUTH, PIPE_NORTH_EAST, PIPE_NORTH_WEST}
	PIPES_EAST  = []byte{PIPE_EAST_WEST, PIPE_SOUTH_EAST, PIPE_NORTH_EAST}
	PIPES_WEST  = []byte{PIPE_EAST_WEST, PIPE_SOUTH_WEST, PIPE_NORTH_WEST}
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	var array [][]byte

	for _, line := range lines {
		array = append(array, []byte(line))
	}

	field := fields.NewFromArray(array)
	startPosition, err := field.FindPosition(func(cell byte) bool {
		return cell == PIPE_START
	})
	if err != nil {
		log.Fatalln("Cannot find start position")
	}

	var positions []fields.FieldPosition
	visited := make(map[fields.FieldPosition]bool)
	visited[startPosition] = true

	if next, err := moveNorth(field, visited, startPosition); err == nil {
		positions = append(positions, next)
	}

	if next, err := moveSouth(field, visited, startPosition); err == nil {
		positions = append(positions, next)
	}

	if next, err := moveEast(field, visited, startPosition); err == nil {
		positions = append(positions, next)
	}

	if next, err := moveWest(field, visited, startPosition); err == nil {
		positions = append(positions, next)
	}

	forward := positions[0]
	visited[forward] = true
	backward := positions[1]
	visited[backward] = true
	steps := 1

	for !forward.Equals(backward) && !forward.Neighbors(backward) {
		// fmt.Println(forward, backward)
		forward, err = move(field, visited, forward)
		if err != nil {
			log.Fatalln("Cannot move forward", err)
		}
		backward, err = move(field, visited, backward)
		if err != nil {
			log.Fatalln("Cannot move backward", err)
		}

		visited[forward] = true
		visited[backward] = true
		steps += 1
	}

	cleanField := fields.New(field.Width, field.Height, GROUND)
	for position, isVisited := range visited {
		if isVisited {
			cleanField.Set(position, field.Get(position))
		}
	}

	insideArea := 0

	for y, row := range cleanField.Cells {
		var lastBend byte
		var isInside bool

		for x, cell := range row {
			position := fields.FieldPosition{X: x, Y: y}
			switch cell {
			case GROUND:
				if isInside {
					insideArea += 1
					cleanField.Set(position, 'I')
				}
			case PIPE_NORTH_SOUTH:
				isInside = !isInside
			case PIPE_NORTH_EAST:
				lastBend = PIPE_NORTH_EAST
			case PIPE_SOUTH_EAST:
				lastBend = PIPE_SOUTH_EAST
			case PIPE_SOUTH_WEST:
				if lastBend == PIPE_NORTH_EAST {
					isInside = !isInside
				} else if lastBend == PIPE_SOUTH_EAST {
					lastBend = 0
				}
			case PIPE_NORTH_WEST:
				if lastBend == PIPE_NORTH_EAST {
					lastBend = 0
				} else if lastBend == PIPE_SOUTH_EAST {
					isInside = !isInside
				}
			}
		}
	}

	for _, row := range cleanField.Cells {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}

	fmt.Println("Furthest point:", steps)
	fmt.Println("Area:", insideArea)
}

func moveNorth(field *fields.Field, visited map[fields.FieldPosition]bool, position fields.FieldPosition) (fields.FieldPosition, error) {
	if next := position.NextNorth(); field.Available(next) && !visited[next] && slices.Contains(PIPES_SOUTH, field.Get(next)) {
		return next, nil
	}
	return fields.FieldPosition{}, errors.New("cannot move north")
}

func moveSouth(field *fields.Field, visited map[fields.FieldPosition]bool, position fields.FieldPosition) (fields.FieldPosition, error) {
	if next := position.NextSouth(); field.Available(next) && !visited[next] && slices.Contains(PIPES_NORTH, field.Get(next)) {
		return next, nil
	}
	return fields.FieldPosition{}, errors.New("cannot move south")
}

func moveEast(field *fields.Field, visited map[fields.FieldPosition]bool, position fields.FieldPosition) (fields.FieldPosition, error) {
	if next := position.NextEast(); field.Available(next) && !visited[next] && slices.Contains(PIPES_WEST, field.Get(next)) {
		return next, nil
	}
	return fields.FieldPosition{}, errors.New("cannot move east")
}

func moveWest(field *fields.Field, visited map[fields.FieldPosition]bool, position fields.FieldPosition) (fields.FieldPosition, error) {
	if next := position.NextWest(); field.Available(next) && !visited[next] && slices.Contains(PIPES_EAST, field.Get(next)) {
		return next, nil
	}
	return fields.FieldPosition{}, errors.New("cannot move west")
}

func move(field *fields.Field, visited map[fields.FieldPosition]bool, position fields.FieldPosition) (fields.FieldPosition, error) {

	if slices.Contains(PIPES_NORTH, field.Get(position)) {
		if next, err := moveNorth(field, visited, position); err == nil {
			return next, nil
		}
	}

	if slices.Contains(PIPES_SOUTH, field.Get(position)) {
		if next, err := moveSouth(field, visited, position); err == nil {
			return next, nil
		}
	}

	if slices.Contains(PIPES_EAST, field.Get(position)) {
		if next, err := moveEast(field, visited, position); err == nil {
			return next, nil
		}
	}

	if slices.Contains(PIPES_WEST, field.Get(position)) {
		if next, err := moveWest(field, visited, position); err == nil {
			return next, nil
		}
	}

	return fields.FieldPosition{}, errors.New("cannot move")
}
