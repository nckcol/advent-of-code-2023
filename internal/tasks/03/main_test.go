package main

import (
	"testing"
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
