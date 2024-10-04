package collection

import (
	"fmt"
	"iter"
	"slices"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return &s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Iterator() iter.Seq[T] {
	return slices.Values(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}
