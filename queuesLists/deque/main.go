package deque

type Deque[T any] struct {
	items []T
}

// insertFront
func (d *Deque[T]) InsertFront(item T) {
	d.items = append(d.items, item)
	for i := len(d.items) - 1; i > 0; i-- {
		d.items[i] = d.items[i-1]
	}
	d.items[0] = item
}

// InsertBack
func (d *Deque[T]) InsertBack(item T) {
	d.items = append(d.items, item)
}

// first
func (d *Deque[T]) First() T {
	return d.items[0]
}

// RemoveFirst
func (d *Deque[T]) RemoveFirst() T {
	returnValue := d.items[0]
	d.items = d.items[1:]
	return returnValue
}

// Last
func (d *Deque[T]) Last() T {
	length := len(d.items)
	return d.items[length-1]
}

// RemoveLast
func (d *Deque[T]) RemoveLast() T {
	length := len(d.items)
	returnValue := d.items[length-1]
	d.items = d.items[:length-1]
	return returnValue
}

// Empty
func (d *Deque[T]) Empty() bool {
	return len(d.items) == 0
}
