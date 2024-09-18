package doublylinkedlist

import "fmt"

type Ordered interface {
	~int | ~string | ~float64
}

type Node[T Ordered] struct {
	Item T
	next *Node[T]
	prev *Node[T]
}
type List[T Ordered] struct {
	first       *Node[T]
	last        *Node[T]
	numberItems int
}

// Methods
func (list *List[T]) Append(item T) {
	newNode := Node[T]{item, nil, nil}
	if list.first == nil {
		list.first = &newNode
		list.last = list.first
	} else {
		list.last.next = &newNode
		newNode.prev = list.last
		list.first.prev = &newNode
		newNode.next = list.first
	}
	list.numberItems++
}

func (list *List[T]) InsertAt(index int, item T) error {
	if index < 0 || index > list.numberItems {
		return fmt.Errorf("invalid index")
	}
	newNode := Node[T]{item, nil, nil}
	if index == 0 {
		newNode.next = list.first
		if list.first != nil {
			list.first.prev = &newNode
		}
		list.first = &newNode
		list.numberItems += 1
		if list.numberItems == 1 {
			list.last = list.first
		}
		return nil
	}
	count := 0
	node := list.first
	previous := node
	for count < index {
		previous = node
		node = node.next
		count++
	}
	newNode.next = node
	newNode.prev = previous
	previous.next = &newNode
	node.prev = &newNode
	list.numberItems++
	return nil
}

func (list *List[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero, fmt.Errorf("invalid index")
	}
	node := list.first
	if index == 0 {
		toRemove := node
		list.first = toRemove.next
		list.numberItems--
		if list.numberItems == 1 {
			list.last = list.first
		}
		return toRemove.Item, nil
	}
	count := 0
	previous := node
	for count < list.numberItems {
		node = node.next
		previous = node
		count++
	}
	toRemove := node
	previous.next = toRemove.next
	toRemove.next.prev = previous
	list.numberItems--
	if list.numberItems <= 1 {
		list.last = list.first
	}
	return toRemove.Item, nil
}

func (list *List[T]) IndexOf(item T) int {
	node := list.first
	index := 0
	for {
		if node.Item == item {
			return index
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		index++
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
			break
		}
		node = node.next
	}
	return node.next.Item
}

func (list *List[T]) ItemBefore(item T) T {
	node := list.first
	for {
		if node == nil {
			var zero T
			return zero
		}

		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.prev.Item
}

func (list *List[T]) Items() []T {
	items := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		items = append(items, node.Item)
		node = node.next
	}
	return items
}

func (list *List[T]) ReverseItems() []T {
	results := []T{}
	node := list.last
	for {
		if node == nil {
			break
		}
		results = append(results, node.Item)
		node = node.prev
	}
	return results
}

// first
func (list *List[T]) First() T {
	if list.first == nil {
		var zero T
		return zero
	}
	return list.first.Item
}

// last
func (list *List[T]) Last() T {
	if list.first == nil {
		var zero T
		return zero
	}
	return list.last.Item
}

// size
func (list *List[T]) Size() int {
	return list.numberItems
}
