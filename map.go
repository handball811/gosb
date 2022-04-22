package main

type Slice[T any] struct {
	s []T
}

func NewSlice[T any](size int) *Slice[T] {
	return &Slice[T]{
		s: make([]T, size),
	}
}

func From[T any](s []T) *Slice[T] {
	return &Slice[T]{
		s: s,
	}
}

func (s *Slice[T]) S() []T {
	return s.s
}
