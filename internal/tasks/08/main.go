package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nckcol/advent-of-code-2023/internal/utils/graph"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
	"github.com/nckcol/advent-of-code-2023/internal/utils/numbers"
)

const (
	DIRECTION_LEFT  = 'L'
	DIRECTION_RIGHT = 'R'
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	directions := []byte(lines[0])
	nodeLookup := make(graph.Lookup[string], 0)

	for _, line := range lines[2:] {
		name := line[0:3]
		leftName := line[7:10]
		rightName := line[12:15]

		node := nodeLookup.Ensure(name)
		node.Left = nodeLookup.Ensure(leftName)
		node.Right = nodeLookup.Ensure(rightName)
	}

	nodes := make([]*graph.Node[string], 0)
	for _, node := range nodeLookup {
		if strings.HasSuffix(node.Value, "A") {
			nodes = append(nodes, node)
		}
	}
	steps := make([]int, len(nodes))

	for i, n := range nodes {
		steps[i] = 0
		node := n
		for node != nil && !strings.HasSuffix(node.Value, "Z") {
			direction := directions[steps[i]%len(directions)]
			if direction == DIRECTION_LEFT {
				node = node.Left
			} else if direction == DIRECTION_RIGHT {
				node = node.Right
			} else {
				log.Fatalf("Unknown direction: %x", direction)
			}
			steps[i] += 1
		}
	}

	fmt.Println("Done in", numbers.LcmSlice(steps), "steps")
}
