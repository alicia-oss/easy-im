package timing_queue

import (
	"container/heap"
	"time"
)

func NewTimeQueue() *TimingQueue {
	return &TimingQueue{queue: &priorityQueue{}}
}

type TimingQueue struct {
	queue *priorityQueue
}

func (t *TimingQueue) Push(tm time.Time, value interface{}) {
	itemT := &item{
		value:    value,
		priority: tm,
	}
	heap.Push(t.queue, itemT)
}

func (t *TimingQueue) Pop() (time.Time, interface{}) {
	if t.queue.Len() == 0 {
		return time.Now(), nil
	}
	itemT := heap.Pop(t.queue).(*item)
	return itemT.priority, itemT.value
}
func (t *TimingQueue) Peek() (time.Time, interface{}) {
	if t.queue.Len() == 0 {
		return time.Now(), nil
	}
	itemT := t.queue.Peek().(*item)
	return itemT.priority, itemT.value
}
