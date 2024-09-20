package set

type Ordered interface {
	~int | ~float64 | ~string
}

type Set[T Ordered] struct {
	items map[T]bool
}

// Insert
func (s *Set[T]) Insert(item T) {
	if s.items == nil {
		s.items = make(map[T]bool)
	}
	// Prevent duplicate entry
	_, present := s.items[item]
	if !present {
		s.items[item] = true
	}
}

// Delete
func (s *Set[T]) Delete(item T) {
	_, present := s.items[item]
	if present {
		delete(s.items, item)
	}
}

// In
func (s *Set[T]) In(item T) bool {
	_, present := s.items[item]
	return present
}

// Items
func (s *Set[T]) Items() []T {
	results := []T{}
	for i := range s.items {
		results = append(results, i)
	}
	return results
}

// Size
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Union
func (s *Set[T]) Union(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for index := range s.items {
		result.items[index] = true
	}
	for index := range set2.items {
		_, present := result.items[index]
		if !present {
			result.items[index] = true
		}
	}
	return &result
}

// Intersection
func (s *Set[T]) Intersection(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for index := range set2.items {
		_, present := s.items[index]
		if present {
			result.items[index] = true
		}
	}
	return &result
}

// Difference
func (s *Set[T]) Difference(set2 Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for index := range s.items {
		_, present := set2.items[index]
		if !present {
			result.items[index] = true
		}
	}
	return &result
}

// Subset
func (s *Set[T]) Subset(set2 Set[T]) bool {
	for index := range set2.items {
		_, present := s.items[index]
		if !present {
			return false
		}
	}
	return true
}
