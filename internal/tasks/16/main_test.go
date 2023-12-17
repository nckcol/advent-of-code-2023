package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := calculateDistance(Galaxy{X: 1, Y: 6}, Galaxy{X: 5, Y: 11})
		assert.Equal(t, 9, got)
	})
}
