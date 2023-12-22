package graph

type NodeN[T any] struct {
	Value    T
	Children []*NodeN[T]
}

func (node *NodeN[T]) Find(fn func(node *NodeN[T]) bool) *NodeN[T] {
	stack := []*NodeN[T]{node}

	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]
		if fn(n) {
			return n
		}

		stack = append(stack, n.Children...)
	}

	return nil
}
