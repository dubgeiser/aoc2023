package collections

type Set[T comparable] struct {
	// Empty struct uses no memory
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{make(map[T]struct{})}
	return s
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) AddSet(s2 *Set[T]) {
	for t := range s2.items {
		s.Add(t)
	}
}

func (s *Set[T]) Remove(v T) {
	delete(s.items, v)
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Items() map[T]struct{} {
	// Iterating over keys of a map:
	// for key := range map {...}
	// Effectively looping over the items in a set
	return s.items
}
