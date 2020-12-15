package stack

import (
	"sync"
)

type StackMu struct {
	sync.Mutex
	head *nodeMu
}

type nodeMu struct {
	value interface{}
	next  *nodeMu
}

func (s *StackMu) Push(v interface{}) {
	node := &nodeMu{value: v, next: nil}
	s.Lock()
	if s.head != nil {
		node.next = s.head
	}
	s.head = node
	s.Unlock()
}

func (s *StackMu) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	if s.head == nil {
		return nil
	}
	node := *s.head
	s.head = s.head.next
	return node.value
}
