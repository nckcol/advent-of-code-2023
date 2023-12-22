package graph

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (node *Node[T]) Find(fn func(node *Node[T]) bool) *Node[T] {
	stack := []*Node[T]{node}

	for len(stack) > 0 {
		n := stack[0]
		stack = stack[1:]
		if fn(n) {
			return n
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}

	return nil
}
