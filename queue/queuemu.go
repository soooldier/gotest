package queue

import "sync"

type queueMu struct {
	sync.Mutex
	head *nodeMu
	tail *nodeMu
}

type nodeMu struct {
	value interface{}
	next  *nodeMu
}

func NewQueueMu() *queueMu {
	dummy := &nodeMu{next: nil}
	return &queueMu{head: dummy}
}

func (q *queueMu) Lpush(v interface{}) {
	node := &nodeMu{value: v}
	q.Lock()
	defer q.Unlock()
	if q.head.next == nil {
		q.head.next = node
		return
	}
	node.next = q.head.next
	q.head.next = node
}

func (q *queueMu) Rpop() interface{} {
	q.Lock()
	defer q.Unlock()
	old := q.head.next
	if old == nil {
		return nil
	}
	q.head.next = old.next
	return old
}
