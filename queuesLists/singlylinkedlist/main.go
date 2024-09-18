package singlylinkedlist

import "fmt"

type Ordered interface {
	~int | ~string | ~float64
}

type Node[T Ordered] struct {
	Item T
	next *Node[T]
}

type List[T Ordered] struct {
	first       *Node[T]
	numberItems int
}

// append
func (list *List[T]) Append(item T) {
	newNode := Node[T]{item, nil}
	if list.first == nil {
		list.first = &newNode
	} else {
		last := list.first
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &newNode
	}
	list.numberItems += 1
}

func (list *List[T]) InsertAt(index int, item T) error {
	if index < 0 || index > list.numberItems {
		return fmt.Errorf("index out of bound")
	}

	newNode := Node[T]{item, nil}
	if index == 0 {
		newNode.next = list.first
		list.first = &newNode
		list.numberItems++
		return nil
	}

	current := list.first
	count := 0
	previous := current
	for count < index {
		previous = current
		current = current.next
		count++
	}
	newNode.next = current
	previous.next = &newNode
	list.numberItems++
	return nil
}

func (list *List[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero , fmt.Errorf("index is out of bounds")
	}

	if index == 0 {
		output :=  list.first
		list.first = output.next
		list.numberItems--
		return output.Item, nil
	}

	node := list.first
	previous := node
	count := 0
	for count < index {
		previous = node
		count++
		node = node.next
	}
	toRemove := node
	previous.next = toRemove.next
	list.numberItems -= 1
	return toRemove.Item, nil
}

func (list *List[T]) IndexOf(item T) int {
	var zero T
	if item == zero {
		return -1
	}
	node := list.first
	count := 0
	for {
		if node.Item == item {
			return count
		}
		if node.next == nil {
			return	-1
		}
		node = node.next
		count++
	}
}

func (list *List[T]) ItemAfter(item T) T {
	node := list.first
	for {
		if node == nil {
			var zero T
			return zero
		}
		if node.Item == item {
			return node.next.Item
		}
		node = node.next
	}
}

func (list *List[T]) Items() []T{
	result := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		result = append(result, node.Item)
		node = node.next
	}
	return result
}

// First
func (list *List[T]) First() *Node[T] {
	return list.first
}


// Size
func (list *List[T]) Size() int {
	return list.numberItems
}
