package nodequeue

type Node[T any] struct {
	item T
	next *Node[T]
}

type Queue[T any] struct {
	first, last *Node[T]
	length      int
}

type Iterator[T any] struct {
	next *Node[T]
}

// Methods
func (q *Queue[T]) Insert(item T) {
	newNode := &Node[T]{item, nil}
	if q.first == nil {
		q.first = newNode
		q.last = q.first
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length += 1
}

func (q *Queue[T]) Remove() T {
	if q.first == nil {
		var zeroValue T
		return zeroValue
	}
	returnValue := q.first.item
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	q.length--
	return returnValue
}

func (q Queue[T]) First() T {
	return q.first.item
}

func (q Queue[T]) Size() int {
	return q.length
}

func (q *Queue[T]) Range() Iterator[T] {
	return Iterator[T]{q.first}
}

func (i *Iterator[T]) Empty() bool {
	return i.next == nil
}

func (i *Iterator[T]) Next() T {
	returnValue := i.next.item
	if i.next != nil {
		i.next = i.next.next
	}
	return returnValue
}
