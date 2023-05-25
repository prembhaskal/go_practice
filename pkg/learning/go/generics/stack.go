package generics

type Stack[T any] struct {
	ar []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		ar: make([]T, 0),
	}
}

func (s *Stack[T]) push(item T) {
	s.ar =   append(s.ar, item)
}

func (s *Stack[T]) pop() T {
	var item T
	if s.isempty() {
		return item
	}
	item = s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return item
}

func (s *Stack[T]) isempty() bool {
	return len(s.ar) == 0}
