package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateGamePower(t *testing.T) {
	t.Run("Card 1", func(t *testing.T) {
		card := parseCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
		value := calculateCardValue(card)

		assert.Equal(t, 8, value)
	})
}

func BenchmarkCalculateGamePower(b *testing.B) {
	source := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	for n, value := range source {
		b.Run(fmt.Sprintf("value_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				card := parseCard(value)
				calculateCardValue(card)
			}
		})
	}
}
