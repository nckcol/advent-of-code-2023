package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestGetGears(t *testing.T) {
	t.Run("should count numbers for every gear", func(t *testing.T) {
		engineMap := parseEngineMap(
			[]string{
				"1*114*.",
				".*.*...",
			},
		)
		gears := getGears(engineMap)

		sum := 0
		for _, gear := range gears {
			sum += gear.Node1.Number * gear.Node2.Number
		}

		assert.Equal(t, 228, sum)
	})
}

func BenchmarkGetPartNumbers(b *testing.B) {
	b.Run("small map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			engineMap := parseEngineMap(input)
			getPartNumbers(engineMap)
		}
	})
}

func BenchmarkGetGears(b *testing.B) {
	b.Run("small map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			engineMap := parseEngineMap(input)
			getGears(engineMap)
		}
	})
}
