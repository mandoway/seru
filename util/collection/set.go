package collection

import (
	"fmt"
	"maps"
	"slices"
)

type Set struct {
	items map[int]struct{}
}

func NewSet(init []int) *Set {
	set := &Set{items: make(map[int]struct{})}
	for _, item := range init {
		set.Put(item)
	}
	return set
}

func (s *Set) Put(item int) {
	s.items[item] = struct{}{}
}

func (s *Set) Has(item int) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) String() string {
	return fmt.Sprint(slices.Collect(maps.Keys(s.items)))
}
