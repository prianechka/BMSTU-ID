package huffman

import (
	"sync"
)

type Item uint32

type Stack struct {
	items []Item
	lock  sync.RWMutex
}

func (s *Stack) New() *Stack {
	s.items = []Item{}
	return s
}

func (s *Stack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *Stack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}
