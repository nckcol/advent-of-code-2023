package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateGamePower(t *testing.T) {
	t.Run("Game 1", func(t *testing.T) {
		gameNumber, gamePower := calculateGamePower("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

		assert.Equal(t, 1, gameNumber)
		assert.Equal(t, 48, gamePower)
	})

	t.Run("Game 2", func(t *testing.T) {
		gameNumber, gamePower := calculateGamePower("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")

		assert.Equal(t, 2, gameNumber)
		assert.Equal(t, 12, gamePower)
	})

	t.Run("Game 3", func(t *testing.T) {
		gameNumber, gamePower := calculateGamePower("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")

		assert.Equal(t, 3, gameNumber)
		assert.Equal(t, 1560, gamePower)
	})

	t.Run("Game 4", func(t *testing.T) {
		gameNumber, gamePower := calculateGamePower("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")

		assert.Equal(t, 4, gameNumber)
		assert.Equal(t, 630, gamePower)
	})

	t.Run("Game 5", func(t *testing.T) {
		gameNumber, gamePower := calculateGamePower("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

		assert.Equal(t, 5, gameNumber)
		assert.Equal(t, 36, gamePower)
	})
}

func BenchmarkCalculateGamePower(b *testing.B) {
	source := []string{
		"Game 12: 3 red, 2 green, 15 blue; 1 blue, 1 green, 4 red; 1 green, 12 blue, 3 red; 1 red, 10 blue; 3 red, 2 green, 14 blue; 3 red, 13 blue",
		"Game 13: 7 blue, 5 red; 7 red, 3 green, 9 blue; 9 green, 7 blue, 7 red; 6 blue, 8 red; 11 red; 3 green, 7 blue, 8 red",
		"Game 14: 4 blue, 6 green, 7 red; 8 red, 4 green, 11 blue; 3 green, 9 red, 13 blue",
		"Game 15: 3 green, 1 blue, 5 red; 2 red; 1 red, 4 green",
		"Game 16: 1 green, 7 blue; 3 red, 5 blue; 1 green, 5 blue; 5 blue, 1 green; 1 green, 1 red, 13 blue",
		"Game 17: 4 blue, 2 red, 4 green; 1 blue, 7 red, 4 green; 4 red, 4 green, 10 blue; 1 blue, 4 red, 14 green",
		"Game 18: 7 blue, 5 green; 4 blue, 3 green; 1 red, 6 green, 7 blue",
		"Game 19: 10 blue, 3 red, 6 green; 3 blue, 4 red, 17 green; 19 green, 3 red, 3 blue; 19 green, 3 blue; 4 red, 7 green, 7 blue; 10 blue, 13 green, 1 red",
		"Game 20: 3 blue, 6 red, 1 green; 6 green, 7 red, 18 blue; 1 green, 5 red, 14 blue; 1 green, 12 blue, 8 red",
		"Game 21: 16 blue, 7 green, 13 red; 11 red, 7 blue, 5 green; 4 green, 3 blue",
	}

	for n, value := range source {
		b.Run(fmt.Sprintf("value_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				calculateGamePower(value)
			}
		})
	}
}
