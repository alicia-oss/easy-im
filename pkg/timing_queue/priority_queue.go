package timing_queue

import (
	"container/heap"
	"time"
)

type item struct {
	value    interface{} // notice id
	priority time.Time   // 优先级队列中节点的优先级
	index    int         // index是该节点在堆中的位置
}

// priorityQueue 优先级队列需要实现heap的interface
type priorityQueue []*item

// Len 绑定Len方法
func (pq priorityQueue) Len() int {
	return len(pq)
}

// Less 绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority.Before(pq[j].priority)
}

// Swap 绑定swap方法
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

// Pop 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *priorityQueue) Peek() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	return item
}

// Push 绑定push方法
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *priorityQueue) update(item *item, value uint64, priority time.Time) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
