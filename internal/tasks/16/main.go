package main

import (
	"fmt"
	"log"

	"github.com/nckcol/advent-of-code-2023/internal/utils/fields"
	"github.com/nckcol/advent-of-code-2023/internal/utils/graph"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	CELL_MIRROR_F   = '/'
	CELL_MIRROR_B   = '\\'
	CELL_SPLITTER_V = '|'
	CELL_SPLITTER_H = '-'
	CELL_EMPTY      = '.'
	CELL_ENERGIZED  = '#'
)

type Direction int

const (
	DIRECTION_NONE  Direction = 0
	DIRECTION_RIGHT Direction = 1
	DIRECTION_UP    Direction = 2
	DIRECTION_LEFT  Direction = 3
	DIRECTION_DOWN  Direction = 4
)

type Beam struct {
	Position  fields.FieldPosition
	Direction Direction
}

type Trace struct {
	Position  fields.FieldPosition
	Direction Direction
}

type Path []Trace

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	field := fields.NewFromStringArray(lines)

	beams := make([]Beam, 0)
	for x := 0; x < field.Width; x++ {
		beams = append(beams,
			Beam{Position: fields.FieldPosition{X: x, Y: -1}, Direction: DIRECTION_DOWN},
			Beam{Position: fields.FieldPosition{X: x, Y: field.Height}, Direction: DIRECTION_UP},
		)
	}
	for y := 0; y < field.Height; y++ {
		beams = append(beams,
			Beam{Position: fields.FieldPosition{X: -1, Y: y}, Direction: DIRECTION_RIGHT},
			Beam{Position: fields.FieldPosition{X: field.Width, Y: y}, Direction: DIRECTION_LEFT},
		)
	}

	max := 0

	for _, beam := range beams {
		traceLookup := map[Trace]bool{}
		path := traverse(field, beam, traceLookup)

		fmt.Println("Done traversing")

		energizedField := fields.New(field.Width, field.Height, CELL_EMPTY)
		energize(energizedField, path)
		// fmt.Println(energizedField)

		energized := energizedField.FindAllPositions(func(cell byte) bool {
			return cell == CELL_ENERGIZED
		})

		if len(energized) > max {
			max = len(energized)
		}
		fmt.Println("Energized cells:", len(energized))
	}

	fmt.Println("Max energized cells:", max)
}

func traverse(field *fields.Field, beam Beam, traceLookup map[Trace]bool) *graph.Node[Trace] {
	if traceLookup[getBeamTrace(beam)] {
		return &graph.Node[Trace]{Value: Trace{Position: beam.Position, Direction: DIRECTION_NONE}, Left: nil, Right: nil}
	}

	traceLookup[getBeamTrace(beam)] = true
	node := &graph.Node[Trace]{Value: getBeamTrace(beam), Left: nil, Right: nil}

	for {
		beam.Position = nextPosition(beam.Position, beam.Direction)
		if !field.Available(beam.Position) {
			return node
		} else if field.Get(beam.Position) != CELL_EMPTY {
			break
		}
	}

	directions := nextDirections(field.Get(beam.Position), beam.Direction)
	if len(directions) >= 1 {
		node.Left = traverse(field, Beam{
			Position:  beam.Position,
			Direction: directions[0],
		}, traceLookup)
	}
	if len(directions) >= 2 {
		node.Right = traverse(field, Beam{
			Position:  beam.Position,
			Direction: directions[1],
		}, traceLookup)
	}
	return node
}

func energize(field *fields.Field, node *graph.Node[Trace]) {
	if node == nil || node.Value.Direction == DIRECTION_NONE {
		return
	}
	position := node.Value.Position
	for {
		if field.Available(position) {
			field.Set(position, CELL_ENERGIZED)
		}
		position = nextPosition(position, node.Value.Direction)
		if !field.Available(position) || (node.Left != nil && position == node.Left.Value.Position) {
			break
		}
	}
	energize(field, node.Left)
	energize(field, node.Right)
}

func nextPosition(position fields.FieldPosition, direction Direction) fields.FieldPosition {
	switch direction {
	case DIRECTION_RIGHT:
		return fields.FieldPosition{X: position.X + 1, Y: position.Y}
	case DIRECTION_UP:
		return fields.FieldPosition{X: position.X, Y: position.Y - 1}
	case DIRECTION_LEFT:
		return fields.FieldPosition{X: position.X - 1, Y: position.Y}
	case DIRECTION_DOWN:
		return fields.FieldPosition{X: position.X, Y: position.Y + 1}
	}
	return position
}

func nextDirections(cell byte, direction Direction) []Direction {
	switch cell {
	case CELL_MIRROR_F:
		switch direction {
		case DIRECTION_RIGHT:
			return []Direction{DIRECTION_UP}
		case DIRECTION_UP:
			return []Direction{DIRECTION_RIGHT}
		case DIRECTION_LEFT:
			return []Direction{DIRECTION_DOWN}
		case DIRECTION_DOWN:
			return []Direction{DIRECTION_LEFT}
		}
	case CELL_MIRROR_B:
		switch direction {
		case DIRECTION_RIGHT:
			return []Direction{DIRECTION_DOWN}
		case DIRECTION_UP:
			return []Direction{DIRECTION_LEFT}
		case DIRECTION_LEFT:
			return []Direction{DIRECTION_UP}
		case DIRECTION_DOWN:
			return []Direction{DIRECTION_RIGHT}
		}
	case CELL_SPLITTER_V:
		switch direction {
		case DIRECTION_RIGHT, DIRECTION_LEFT:
			return []Direction{DIRECTION_UP, DIRECTION_DOWN}
		case DIRECTION_DOWN, DIRECTION_UP:
			return []Direction{direction}
		}
	case CELL_SPLITTER_H:
		switch direction {
		case DIRECTION_RIGHT, DIRECTION_LEFT:
			return []Direction{direction}
		case DIRECTION_UP, DIRECTION_DOWN:
			return []Direction{DIRECTION_LEFT, DIRECTION_RIGHT}
		}
	}
	return []Direction{}
}

func getBeamTrace(beam Beam) Trace {
	return Trace(beam)
}

func (d Direction) String() string {
	switch d {
	case DIRECTION_RIGHT:
		return "R"
	case DIRECTION_UP:
		return "U"
	case DIRECTION_LEFT:
		return "L"
	case DIRECTION_DOWN:
		return "D"
	}
	return ""
}
