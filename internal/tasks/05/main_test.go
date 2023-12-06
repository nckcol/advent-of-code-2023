package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRange(t *testing.T) {
	t.Run("Range 1", func(t *testing.T) {
		parsedRange, err := parseRange("0 15 37")

		assert.NoError(t, err)
		assert.Equal(t, RangeMap{
			DestinationStart: 0,
			SourceStart:      15,
			Size:             37,
		}, parsedRange)
	})

	t.Run("Range 2", func(t *testing.T) {
		parsedRange, err := parseRange("3828940251 2532239941 80616634")

		assert.NoError(t, err)
		assert.Equal(t, RangeMap{
			DestinationStart: 3828940251,
			SourceStart:      2532239941,
			Size:             80616634,
		}, parsedRange)
	})
}

func TestMapValues(t *testing.T) {
	t.Run("seed", func(t *testing.T) {
		seed := []Range{
			{Start: 13, Size: 1},
			{Start: 14, Size: 1},
			{Start: 55, Size: 1},
			{Start: 79, Size: 1},
		}
		mapped := mapValues(seed, []RangeMap{
			{
				SourceStart:      50,
				DestinationStart: 52,
				Size:             48,
			},
			{
				SourceStart:      98,
				DestinationStart: 50,
				Size:             2,
			},
		})

		assert.Equal(t, []Range{
			{Start: 13, Size: 1},
			{Start: 14, Size: 1},
			{Start: 57, Size: 1},
			{Start: 81, Size: 1},
		}, mapped)
	})

	t.Run("soil", func(t *testing.T) {
		soil := []Range{
			{Start: 13, Size: 1},
			{Start: 14, Size: 1},
			{Start: 57, Size: 1},
			{Start: 81, Size: 1},
		}

		mapped := mapValues(soil, []RangeMap{
			{
				SourceStart:      0,
				DestinationStart: 39,
				Size:             15,
			},
			{
				SourceStart:      15,
				DestinationStart: 0,
				Size:             37,
			},
			{
				SourceStart:      52,
				DestinationStart: 37,
				Size:             2,
			},
		})

		assert.Equal(t, []Range{
			{Start: 52, Size: 1},
			{Start: 53, Size: 1},
			{Start: 57, Size: 1},
			{Start: 81, Size: 1},
		}, mapped)
	})

	t.Run("fertilizer", func(t *testing.T) {
		fertilizer := []Range{
			{Start: 52, Size: 1},
			{Start: 53, Size: 1},
			{Start: 57, Size: 1},
			{Start: 81, Size: 1},
		}

		mapped := mapValues(fertilizer, []RangeMap{
			{
				SourceStart:      0,
				DestinationStart: 42,
				Size:             7,
			},
			{
				SourceStart:      7,
				DestinationStart: 57,
				Size:             4,
			},
			{
				SourceStart:      11,
				DestinationStart: 0,
				Size:             42,
			},
			{
				SourceStart:      53,
				DestinationStart: 49,
				Size:             8,
			},
		})

		assert.Equal(t, []Range{
			{Start: 41, Size: 1},
			{Start: 49, Size: 1},
			{Start: 53, Size: 1},
			{Start: 81, Size: 1},
		}, mapped)
	})

	t.Run("seed range", func(t *testing.T) {
		seed := []Range{
			{Start: 55, Size: 13},
			{Start: 79, Size: 14},
		}
		mapped := mapValues(seed, []RangeMap{
			{
				SourceStart:      50,
				DestinationStart: 52,
				Size:             48,
			},
			{
				SourceStart:      98,
				DestinationStart: 50,
				Size:             2,
			},
		})

		assert.Equal(t, []Range{
			{Start: 57, Size: 13},
			{Start: 81, Size: 14},
		}, mapped)
	})
}

// func BenchmarkCalculateGamePower(b *testing.B) {
// 	source := []string{
// 		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
// 		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
// 		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
// 		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
// 		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
// 		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
// 	}

// 	for n, value := range source {
// 		b.Run(fmt.Sprintf("value_%d", n), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				card := parseCard(value)
// 				calculateCardValue(card)
// 			}
// 		})
// 	}
// }
