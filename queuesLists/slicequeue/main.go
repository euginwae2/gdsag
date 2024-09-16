package slicequeue

type Queue[T any] struct {
	items []T
}

type Iterator[T any] struct {
	next  int //index in items
	items []T
}

// methods
func (q *Queue[T]) Insert(item T) {
	q.items = append(q.items, item)
}

// Remove
func (q *Queue[T]) Remove() T {
	returnValue := q.items[0]
	q.items = q.items[1:]
	return returnValue
}

func (q Queue[T]) First() T {
	return q.items[0]
}

func (q Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Range() Iterator[T] {
	return Iterator[T]{0, q.items}
}

// Iterator method
func (i *Iterator[T]) Empty() bool {
	return i.next == len(i.items)
}

func (i *Iterator[T]) Next() T {
	returnValue := i.items[i.next]
	i.next++
	return returnValue
}
