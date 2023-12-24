package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
	"github.com/nckcol/advent-of-code-2023/internal/utils/tokenizer"
)

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	Lenses []Lens
}

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := tokenizer.Tokenize(lines[0], tokenizer.WithSeparators([]byte{','}), tokenizer.WithOperators([]string{"=", "-"}))
	if err != nil {
		log.Fatal(err)
	}

	hashCache := make(map[string]byte)
	boxes := make([]Box, 256)
	var (
		box  *Box
		lens Lens = Lens{}
	)
	for _, token := range tokens {
		switch {
		case token.Key == tokenizer.WORD:
			var index byte
			lens.Label = token.Value
			if value, ok := hashCache[lens.Label]; ok {
				index = value
			} else {
				index = hash(lens.Label)
				hashCache[token.Value] = index
			}
			box = &boxes[index]
		case token.Key == tokenizer.NUMBER:
			lens.FocalLength, err = strconv.Atoi(token.Value)
			if err != nil {
				log.Fatal(err)
			}
			box.Ensure(lens)
		case token.Key == tokenizer.OPERATOR && token.Value == "-":
			box.Remove(lens)
		case token.Key == tokenizer.SEPARATOR:
			box = nil
			lens = Lens{}
		}
	}
	printBoxes(boxes)
	fmt.Println("Focusing power:", calculateFocusingPower(boxes))
}

func hash(source string) byte {
	current := 0

	for _, c := range source {
		current += int(c)
		current *= 17
		current %= 256
	}

	return byte(current)
}

func (b *Box) Ensure(lens Lens) {
	for i := range b.Lenses {
		if b.Lenses[i].Label == lens.Label {
			b.Lenses[i].FocalLength = lens.FocalLength
			return
		}
	}
	b.Lenses = append(b.Lenses, lens)
}

func (b *Box) Remove(lens Lens) {
	b.Lenses = slices.DeleteFunc(b.Lenses, func(l Lens) bool {
		return l.Label == lens.Label
	})
}

func (b Box) String() string {
	result := ""
	for _, lens := range b.Lenses {
		result += fmt.Sprintf("[%s %d] ", lens.Label, lens.FocalLength)
	}
	return result
}

func calculateFocusingPower(boxes []Box) int {
	power := 0
	for bi, box := range boxes {
		for li, lens := range box.Lenses {
			power += (bi + 1) * (li + 1) * lens.FocalLength
		}
	}
	return power
}

func printBoxes(boxes []Box) {
	for i, box := range boxes {
		if len(box.Lenses) == 0 {
			continue
		}
		fmt.Printf("Box %v: %v\n", i, box)
	}
}
