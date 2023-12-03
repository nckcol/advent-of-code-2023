package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nckcol/advent-of-code-2023/internal/utils/matrix"
)

type MapNode struct {
	Number int
	Symbol byte
}

func parseEngineMap(input []string) *matrix.Matrix[MapNode] {
	engineMap := matrix.New[MapNode](len(input[0]), len(input))

	for y, line := range input {
		var node *MapNode

		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				node = nil
				continue
			} else if isDigit(line[x]) {
				if node == nil {
					node = &MapNode{
						Number: 0,
					}
				}
				digit := int(line[x] - '0')
				node.Number = node.Number*10 + digit
				engineMap.Set(x, y, node)
			} else {
				node = nil
				engineMap.Set(x, y, &MapNode{
					Symbol: line[x],
				})
			}
		}
	}

	return engineMap
}

func getPartNumbers(engineMap *matrix.Matrix[MapNode]) []*MapNode {
	usedNodes := make(map[*MapNode]bool, 0)
	partNumberNodes := make([]*MapNode, 0)

	for iterator := engineMap.CreateIterator(); iterator.HasNext(); {
		pos, node := iterator.Next()
		x := pos[0]
		y := pos[1]

		if node == nil || node.Symbol == 0 {
			continue
		}

		adjacentNodes := getAdjacentNodes(engineMap, x, y)

		for _, n := range adjacentNodes {
			if n != nil && n.Number != 0 && !usedNodes[n] {
				usedNodes[n] = true
				partNumberNodes = append(partNumberNodes, n)
			}
		}
	}

	return partNumberNodes
}

type Gear struct {
	Node1 *MapNode
	Node2 *MapNode
}

func getGears(engineMap *matrix.Matrix[MapNode]) []Gear {
	gearList := make([]Gear, 0)

	for iterator := engineMap.CreateIterator(); iterator.HasNext(); {
		pos, node := iterator.Next()
		x := pos[0]
		y := pos[1]

		if node == nil || node.Symbol != '*' {
			continue
		}

		adjacentNodes := getAdjacentNodes(engineMap, x, y)
		gear := Gear{}
		isGearValid := false

		for _, n := range adjacentNodes {
			if n == nil || n.Number == 0 {
				continue
			}

			if gear.Node1 == n || gear.Node2 == n {
				continue
			}

			if gear.Node1 == nil {
				gear.Node1 = n
			} else if gear.Node2 == nil {
				gear.Node2 = n
				isGearValid = true
			} else {
				isGearValid = false
			}
		}

		if !isGearValid {
			continue
		}

		gearList = append(gearList, gear)
	}

	return gearList
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func getAdjacentNodes[T any](matrix *matrix.Matrix[T], x int, y int) []*T {
	adjacentNodes := []*T{
		matrix.At(x-1, y-1),
		matrix.At(x, y-1),
		matrix.At(x+1, y-1),
		matrix.At(x-1, y),
		matrix.At(x+1, y),
		matrix.At(x-1, y+1),
		matrix.At(x, y+1),
		matrix.At(x+1, y+1),
	}

	return adjacentNodes
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	engineMap := parseEngineMap(input)

	partNumbers := getPartNumbers(engineMap)
	sum1 := 0
	for _, n := range partNumbers {
		sum1 += n.Number
	}

	gears := getGears(engineMap)
	sum2 := 0
	for _, g := range gears {
		sum2 += g.Node1.Number * g.Node2.Number
	}

	fmt.Println()
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
