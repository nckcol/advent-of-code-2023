package stack

type Stack[T any] struct {
	top    *node[T]
	length int
}

type node[T any] struct {
	value T
	prev  *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	n := &node[T]{value, s.top}
	s.top = n
	s.length++
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Len() == 0 {
		return *new(T), false
	}
	return s.top.value, true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.length == 0 {
		return *new(T), false
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value, true
}

func (s *Stack[T]) Len() int {
	return s.length
}
