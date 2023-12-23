package fields

import (
	"errors"
	"fmt"
)

type Field struct {
	Width  int
	Height int
	Cells  [][]byte
}

func New(width int, height int, fill byte) *Field {
	cells := make([][]byte, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]byte, width)
		if fill != 0 {
			for j := 0; j < width; j++ {
				cells[i][j] = fill
			}
		}
	}
	return &Field{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func NewFromArray(array [][]byte) *Field {
	width := 0
	cells := make([][]byte, len(array))
	for _, row := range array {
		width = max(width, len(row))
	}
	for i, row := range array {
		cells[i] = make([]byte, width)
		copy(cells[i], row)
	}

	return &Field{
		Width:  width,
		Height: len(array),
		Cells:  cells,
	}
}

func NewFromByteArray(array [][]byte) *Field {
	width := 0
	cells := make([][]byte, len(array))
	for _, row := range array {
		width = max(width, len(row))
	}
	for i, row := range array {
		cells[i] = make([]byte, width)
		copy(cells[i], row)
	}

	return &Field{
		Width:  width,
		Height: len(array),
		Cells:  cells,
	}
}

func NewFromStringArray(array []string) *Field {
	width := 0
	cells := make([][]byte, len(array))
	for _, row := range array {
		width = max(width, len(row))
	}
	for i, row := range array {
		cells[i] = make([]byte, width)
		copy(cells[i], row)
	}

	return &Field{
		Width:  width,
		Height: len(array),
		Cells:  cells,
	}
}

func (f *Field) Get(position FieldPosition) byte {
	return f.Cells[position.Y][position.X]
}

func (f *Field) Set(position FieldPosition, value byte) {
	f.Cells[position.Y][position.X] = value
}

func (f *Field) Available(position FieldPosition) bool {
	return position.X >= 0 && position.X < f.Width && position.Y >= 0 && position.Y < f.Height
}

func (m *Field) Col(x int) []byte {
	col := make([]byte, 0, m.Height)
	for y := 0; y < m.Height; y++ {
		col = append(col, m.Get(FieldPosition{x, y}))
	}
	return col
}

func (m *Field) Row(y int) []byte {
	row := make([]byte, 0, m.Width)
	for x := 0; x < m.Width; x++ {
		row = append(row, m.Get(FieldPosition{x, y}))
	}
	return row
}

func (f *Field) FindPosition(predicate func(byte) bool) (FieldPosition, error) {
	for y, row := range f.Cells {
		for x, cell := range row {
			if predicate(cell) {
				return FieldPosition{
					X: x,
					Y: y,
				}, nil
			}
		}
	}
	return FieldPosition{}, errors.New("not found")
}

func (f *Field) FindAllPositions(predicate func(byte) bool) []FieldPosition {
	result := make([]FieldPosition, 0)
	for y, row := range f.Cells {
		for x, cell := range row {
			if predicate(cell) {
				result = append(result, FieldPosition{
					X: x,
					Y: y,
				})
			}
		}
	}
	return result
}

func (f *Field) NextAvailable(position FieldPosition) []FieldPosition {
	result := make([]FieldPosition, 0)
	for _, p := range []FieldPosition{
		position.NextNorth(),
		position.NextSouth(),
		position.NextEast(),
		position.NextWest(),
	} {
		if f.Available(p) {
			result = append(result, p)
		}
	}
	return result
}

func (f *Field) String() string {
	result := ""
	for i, row := range f.Cells {
		if i == len(f.Cells)-1 {
			result += string(row)
		} else {
			result += string(row) + "\n"
		}
	}
	return result
}

type FieldPosition struct {
	X int
	Y int
}

func (p FieldPosition) NextWest() FieldPosition {
	return FieldPosition{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p FieldPosition) NextEast() FieldPosition {
	return FieldPosition{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p FieldPosition) NextNorth() FieldPosition {
	return FieldPosition{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p FieldPosition) NextSouth() FieldPosition {
	return FieldPosition{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p FieldPosition) Equals(other FieldPosition) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p FieldPosition) Neighbors(other FieldPosition) bool {
	return p.NextNorth().Equals(other) || p.NextSouth().Equals(other) || p.NextEast().Equals(other) || p.NextWest().Equals(other)
}

func (p FieldPosition) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p FieldPosition) Move(direction Direction) FieldPosition {
	switch direction {
	case DIRECTION_EAST:
		return p.NextEast()
	case DIRECTION_NORTH:
		return p.NextNorth()
	case DIRECTION_WEST:
		return p.NextWest()
	case DIRECTION_SOUTH:
		return p.NextSouth()
	}
	return p
}
