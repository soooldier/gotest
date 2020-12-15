package stack

import (
	"sync/atomic"
	"unsafe"
)

type Stack struct {
	head unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func (s *Stack) Push(v interface{}) {
	new := &node{value: v}
	for {
		old := load(&s.head)
		new.next = unsafe.Pointer(old)
		if cas(&s.head, old, new) {
			break
		}
	}
}

func (s *Stack) Pop() interface{} {
	for {
		old := load(&s.head)
		if old == nil {
			return nil
		}
		if cas(&s.head, old, load(&old.next)) {
			return old.value
		}
	}
}

func load(p *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
