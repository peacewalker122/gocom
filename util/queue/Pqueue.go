package queue

import (
	"gocom/contract"
	"gocom/util/heap"
)

type QueueType int

const (
	MinQueue QueueType = iota
	MaxQueue QueueType = iota
)

type Queue struct {
	queue contract.Heap
}

func (q *Queue) EnQueue(x interface{}) {
	q.queue.Push(x)
}

func (q *Queue) DeQueue() interface{} {
	return q.queue.Pop()
}

func (q *Queue) Peek() interface{} {
	return q.queue.Top()
}

func (q *Queue) Len() int {
	return q.queue.Len()
}

func (q *Queue) ChangePriority(oldData interface{}, newData interface{}) {
	q.queue.UpdateByValue(oldData, newData)
}

func NewQueue(opt ...QueueType) *Queue {
	var queueType contract.Heap

	queueType = heap.NewMinHeap() // default value

	q := &Queue{
		queue: queueType,
	}
	return q
}
