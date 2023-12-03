package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

type MapNode struct {
	Number int
	Used   bool
	Symbol byte
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

	sum1 := 0
	sum2 := 0

	engineMap := make([][]*MapNode, 1)
	engineMap[0] = make([]*MapNode, 1)

	for _, line := range input {
		engineMap = append(engineMap, make([]*MapNode, 1))
		lastIndex := len(engineMap) - 1
		engineMap[lastIndex] = make([]*MapNode, len(line)+2)
		var node *MapNode

		for i := 0; i < len(line); i++ {
			if line[i] == '.' {
				node = nil
				continue
			} else if isDigit(line[i]) {
				if node == nil {
					node = &MapNode{
						Number: 0,
					}
				}
				digit := int(line[i] - '0')
				node.Number = node.Number*10 + digit
				engineMap[lastIndex][i+1] = node
			} else {
				node = nil
				engineMap[lastIndex][i+1] = &MapNode{
					Symbol: line[i],
				}
			}
		}
	}

	for i, line := range engineMap {
		for j, node := range line {
			if node == nil || node.Symbol == 0 {
				continue
			}
			adjacentNodes := []*MapNode{
				engineMap[i-1][j-1],
				engineMap[i-1][j],
				engineMap[i-1][j+1],
				engineMap[i][j-1],
				engineMap[i][j+1],
				engineMap[i+1][j-1],
				engineMap[i+1][j],
				engineMap[i+1][j+1],
			}

			for _, n := range adjacentNodes {
				if n != nil && n.Number != 0 && !n.Used {
					sum1 += n.Number
					n.Used = true
				}
			}
		}
	}

	for i, line := range engineMap {
		for j, node := range line {
			if node == nil || node.Symbol != '*' {
				continue
			}
			adjacentNodes := []*MapNode{
				engineMap[i-1][j-1],
				engineMap[i-1][j],
				engineMap[i-1][j+1],
				engineMap[i][j-1],
				engineMap[i][j+1],
				engineMap[i+1][j-1],
				engineMap[i+1][j],
				engineMap[i+1][j+1],
			}

			adjacentNumbers := make([]*MapNode, 0)

			for _, n := range adjacentNodes {
				if n != nil && n.Number != 0 {
					used := false
					for _, a := range adjacentNumbers {
						if a == n {
							used = true
							break
						}
					}
					if !used {
						adjacentNumbers = append(adjacentNumbers, n)
					}
				}
			}

			if len(adjacentNumbers) != 2 {
				continue
			}

			sum2 += adjacentNumbers[0].Number * adjacentNumbers[1].Number
		}
	}

	fmt.Println()
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
