package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldHistory(t *testing.T) {
	t.Run("sequence", func(t *testing.T) {
		input := []int{0, 1, 2, 3, 4, 5}
		got := foldHistory(input)
		assert.Equal(t, []int{1, 1, 1, 1, 1}, got)
	})

	t.Run("pair", func(t *testing.T) {
		input := []int{1, 1}
		got := foldHistory(input)
		assert.Equal(t, []int{0}, got)
	})
}

func TestExtrapolateHistoryForwards(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{0, 3, 6, 9, 12, 15}
		slices.Reverse(input)
		got := extrapolateHistoryForwards(input)
		assert.Equal(t, 18, got)
	})
}

func TestExtrapolateHistoryBackwards(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{10, 13, 16, 21, 30, 45}
		got := extrapolateHistoryBackwards(input)
		assert.Equal(t, 5, got)
	})
}
