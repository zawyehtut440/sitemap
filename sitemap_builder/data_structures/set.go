package data_structures

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}

func (s *Set[T]) Contains(val T) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set[T]) Remove(val T) {
	delete(s.m, val)
}
