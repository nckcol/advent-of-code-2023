package graph

type Lookup[T comparable] map[T]*Node[T]

func (l Lookup[T]) Get(value T) *Node[T] {
	return l[value]
}

func (l Lookup[T]) Ensure(value T) *Node[T] {
	if l[value] == nil {
		l[value] = &Node[T]{Value: value}
	}

	return l[value]
}
