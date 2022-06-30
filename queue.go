package lists

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func (queue *Queue[T]) Init() {
	node := &node[T]{}
	cas(&queue.head, nil, node)
	cas(&queue.tail, nil, node)
}

func (queue *Queue[T]) Push(v T) {
	node := &node[T]{
		v: v,
	}
	for {
		tail := load(&queue.tail)
		if cas(&queue.tail, tail, node) {
			tail.next = node
			return
		}
	}
}

func (queue *Queue[T]) Pop(v *T) bool {
	for {
		head := load(&queue.head)
		next := load(&head.next)
		if next == nil {
			return false
		}
		if cas(&queue.head, head, next) {
			*v = next.v
			return true
		}
	}
}
