package matrix

type MatrixNode[T any] struct {
	value *T
}

type Matrix[T any] struct {
	Width  int
	Height int
	items  []MatrixNode[T]
}

func New[T any](width int, height int) *Matrix[T] {
	return &Matrix[T]{
		Width:  width,
		Height: height,
		items:  make([]MatrixNode[T], width*height),
	}
}

func From2DArray[T any](array [][]T) *Matrix[T] {
	height := len(array)
	width := len(array[0])
	for _, row := range array {
		width = max(width, len(row))
	}
	matrix := New[T](width, height)

	for y, row := range array {
		for x, value := range row {
			matrix.Set(x, y, &value)
		}
	}

	return matrix
}

func (m *Matrix[T]) At(x int, y int) *T {
	index := x + y*m.Width
	if index >= len(m.items) {
		return nil
	}
	return m.items[index].value
}

func (m *Matrix[T]) Set(x int, y int, value *T) {
	index := x + y*m.Width
	if index >= len(m.items) {
		return
	}
	m.items[index] = MatrixNode[T]{
		value: value,
	}
}

type MatrixIterator[T any] struct {
	matrix *Matrix[T]
	index  int
}

func (m *Matrix[T]) CreateIterator() *MatrixIterator[T] {
	return &MatrixIterator[T]{
		matrix: m,
		index:  0,
	}
}

func (i *MatrixIterator[T]) HasNext() bool {
	return i.index < i.matrix.Width*i.matrix.Height
}

func (i *MatrixIterator[T]) Next() ([2]int, *T) {
	i.index++
	x := i.index % i.matrix.Width
	y := i.index / i.matrix.Width

	return [2]int{x, y}, i.matrix.At(x, y)
}
