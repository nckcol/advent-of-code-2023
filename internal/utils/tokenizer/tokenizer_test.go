package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	t.Run("accepts white spaces", func(t *testing.T) {
		tokens, _ := Tokenize(" \t  \t\t ")

		assert.Len(t, tokens, 0)
	})

	t.Run("accepts empty string", func(t *testing.T) {
		tokens, _ := Tokenize("")

		assert.Len(t, tokens, 0)
	})

	t.Run("accepts single number", func(t *testing.T) {
		tokens, _ := Tokenize("123")

		assert.Len(t, tokens, 1)
		assert.Equal(t, NUMBER, tokens[0].Key)
		assert.Equal(t, "123", tokens[0].Value)
	})

	t.Run("accepts multiple numbers separated with whitespace", func(t *testing.T) {
		tokens, _ := Tokenize(" 123\t456  \t 789\t")

		assert.Len(t, tokens, 3)
		assert.Equal(t, NUMBER, tokens[0].Key)
		assert.Equal(t, "123", tokens[0].Value)
		assert.Equal(t, NUMBER, tokens[1].Key)
		assert.Equal(t, "456", tokens[1].Value)
		assert.Equal(t, NUMBER, tokens[2].Key)
		assert.Equal(t, "789", tokens[2].Value)
	})

	t.Run("accepts decimal numbers", func(t *testing.T) {
		tokens, _ := Tokenize("3.14159268\t0.12 0.00001")

		assert.Len(t, tokens, 3)
		assert.Equal(t, NUMBER, tokens[0].Key)
		assert.Equal(t, "3.14159268", tokens[0].Value)
		assert.Equal(t, NUMBER, tokens[1].Key)
		assert.Equal(t, "0.12", tokens[1].Value)
		assert.Equal(t, NUMBER, tokens[2].Key)
		assert.Equal(t, "0.00001", tokens[2].Value)
	})

	t.Run("fails to parse wrong decimal numbers", func(t *testing.T) {
		_, err := Tokenize("3..14159268")

		assert.Error(t, ErrNoProgress, err)
	})

	t.Run("accepts single word", func(t *testing.T) {
		tokens, _ := Tokenize("foo")

		assert.Len(t, tokens, 1)
		assert.Equal(t, WORD, tokens[0].Key)
		assert.Equal(t, "foo", tokens[0].Value)
	})

	t.Run("accepts multiple words separated by whitespace", func(t *testing.T) {
		tokens, _ := Tokenize("foo  bar\tBaZ\t\t  \t")

		assert.Len(t, tokens, 3)
		assert.Equal(t, WORD, tokens[0].Key)
		assert.Equal(t, "foo", tokens[0].Value)
		assert.Equal(t, WORD, tokens[1].Key)
		assert.Equal(t, "bar", tokens[1].Value)
		assert.Equal(t, WORD, tokens[2].Key)
		assert.Equal(t, "BaZ", tokens[2].Value)
	})

	t.Run("accepts words and numbers with separators", func(t *testing.T) {
		tokens, _ := Tokenize(
			"Game 1: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
			WithSeparators([]byte{':', ',', ';'}),
		)

		assert.Len(t, tokens, 38)
		assert.Equal(t,
			[]Token{
				{Key: WORD, Value: "Game"},
				{Key: NUMBER, Value: "1"},
				{Key: SEPARATOR, Value: ":"},
				{Key: NUMBER, Value: "1"},
				{Key: WORD, Value: "blue"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "8"},
				{Key: WORD, Value: "green"},
				{Key: SEPARATOR, Value: ";"},
				{Key: NUMBER, Value: "14"},
				{Key: WORD, Value: "green"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "15"},
				{Key: WORD, Value: "blue"},
				{Key: SEPARATOR, Value: ";"},
				{Key: NUMBER, Value: "3"},
				{Key: WORD, Value: "green"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "9"},
				{Key: WORD, Value: "blue"},
				{Key: SEPARATOR, Value: ";"},
				{Key: NUMBER, Value: "8"},
				{Key: WORD, Value: "green"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "8"},
				{Key: WORD, Value: "blue"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "1"},
				{Key: WORD, Value: "red"},
				{Key: SEPARATOR, Value: ";"},
				{Key: NUMBER, Value: "1"},
				{Key: WORD, Value: "red"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "9"},
				{Key: WORD, Value: "green"},
				{Key: SEPARATOR, Value: ","},
				{Key: NUMBER, Value: "10"},
				{Key: WORD, Value: "blue"},
			},
			tokens,
		)
	})
}
