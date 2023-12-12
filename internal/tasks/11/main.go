package main

import (
	"fmt"
	"log"

	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
)

const (
	CELL_GALAXY = '#'
	CELL_SPACE  = '.'
)

type Galaxy struct {
	X int
	Y int
}

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	var galaxies []Galaxy
	galaxiesByX := make(map[int][]*Galaxy)
	galaxiesByY := make(map[int][]*Galaxy)

	xMin := len(lines[0])
	xMax := 0
	yMin := len(lines)
	yMax := 0
	for y, line := range lines {
		for x, cell := range []byte(line) {
			if cell == CELL_GALAXY {
				galaxy := Galaxy{X: x, Y: y}
				galaxies = append(galaxies, galaxy)
				galaxiesByX[x] = append(galaxiesByX[x], &galaxy)
				galaxiesByY[y] = append(galaxiesByY[y], &galaxy)
				xMin = min(xMin, x)
				xMax = max(xMax, x)
				yMin = min(yMin, y)
				yMax = max(yMax, y)
			}
		}
	}

	emptyRows := make([]int, 0)
	for y := yMin; y <= yMax; y++ {
		if len(galaxiesByY[y]) == 0 {
			emptyRows = append(emptyRows, y)
		}
	}

	emptyColumns := make([]int, 0)
	for x := xMin; x <= xMax; x++ {
		if len(galaxiesByX[x]) == 0 {
			emptyColumns = append(emptyColumns, x)
		}
	}

	fmt.Println("Empty rows:", emptyRows)
	fmt.Println("Empty columns:", emptyColumns)

	expansionRate := 1000000
	expandedGalaxies := make([]Galaxy, 0, len(galaxies))
	for _, galaxy := range galaxies {
		expandedCols := 0
		expandedRows := 0
		for expandedCols < len(emptyColumns) && emptyColumns[expandedCols] < galaxy.X {
			expandedCols += 1
		}
		for expandedRows < len(emptyRows) && emptyRows[expandedRows] < galaxy.Y {
			expandedRows += 1
		}
		expandedGalaxies = append(expandedGalaxies, Galaxy{
			X: galaxy.X + expandedCols*(expansionRate-1),
			Y: galaxy.Y + expandedRows*(expansionRate-1),
		})
	}

	sum := 0
	for i := 0; i < len(expandedGalaxies); i++ {
		for j := i + 1; j < len(expandedGalaxies); j++ {
			sum += calculateDistance(expandedGalaxies[i], expandedGalaxies[j])
		}
	}

	fmt.Println("Sum:", sum)
}

func calculateDistance(a, b Galaxy) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
