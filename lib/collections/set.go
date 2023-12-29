package collections

type Set[T comparable] struct {
	// Empty struct uses no memory
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{make(map[T]struct{})}
	return s
}

func NewSetFrom[T comparable] (l []T) *Set[T] {
	s := NewSet[T]()
	s.AddMany(l)
	return s
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) AddMany(l []T) {
	for _, v := range l {
		s.Add(v)
	}
}

func (s *Set[T]) AddSet(s2 *Set[T]) {
	for t := range s2.items {
		s.Add(t)
	}
}

func (s *Set[T]) Remove(v T) {
	delete(s.items, v)
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Export() []T {
	var export []T
	for v := range s.items {
		export = append(export, v)
	}
	return export
}

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	intersection := NewSet[T]()
	for e := range s.items {
		if s2.Has(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func (s *Set[T]) Items() map[T]struct{} {
	// Iterating over keys of a map:
	// for key := range map {...}
	// Effectively looping over the items in a set
	return s.items
}
