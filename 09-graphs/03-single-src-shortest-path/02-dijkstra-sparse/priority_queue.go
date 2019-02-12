package dijkstra

import (
	"container/heap"
	"fmt"
)

type QueueItem struct {
	Priority int // weight of item
	Key      int // key of item
	Payload  interface{}
}

type entry struct {
	QueueItem
	Index int // index of item in queue
}

type priorityQueueImpl struct {
	arr  []*entry
	keys map[int]int // key -> index
}

// heap interface
var _ heap.Interface = (*priorityQueueImpl)(nil)

func newPriorityQueueImpl() *priorityQueueImpl {
	pq := &priorityQueueImpl{
		arr:  make([]*entry, 0),
		keys: make(map[int]int),
	}
	return pq
}

func (pq *priorityQueueImpl) Len() int {
	return len(pq.arr)
}

func (pq *priorityQueueImpl) Less(i, j int) bool {
	return pq.arr[i].Priority < pq.arr[j].Priority
}

func (pq *priorityQueueImpl) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
	pq.arr[i].Index = i
	pq.arr[j].Index = j
	pq.keys[pq.arr[i].Key] = i
	pq.keys[pq.arr[j].Key] = j
}

func (pq *priorityQueueImpl) Pop() interface{} {
	n := pq.Len()
	e := pq.arr[n-1]
	pq.arr = pq.arr[0 : n-1]
	return e.QueueItem
}

func (pq *priorityQueueImpl) Push(x interface{}) {
	n := pq.Len()
	e := &entry{
		QueueItem: x.(QueueItem),
		Index:     n,
	}
	if _, ok := pq.keys[e.Key]; ok {
		panic(fmt.Sprintf("key already exists: %d", e.Key))
	}
	pq.arr = append(pq.arr, e)
	pq.keys[e.Key] = e.Index
}

// PriorityQueue

type PriorityQueue struct {
	impl *priorityQueueImpl
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		impl: newPriorityQueueImpl(),
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return pq.impl.Len()
}

func (pq *PriorityQueue) Enqueue(item QueueItem) {
	// if key exists, update
	impl := pq.impl
	if idx, ok := impl.keys[item.Key]; ok {
		impl.arr[idx].QueueItem = item
		heap.Fix(impl, idx)
	} else {
		heap.Push(impl, item)
	}
}

func (pq *PriorityQueue) Dequeue() QueueItem {
	return heap.Pop(pq.impl).(QueueItem)
}

func (pq *PriorityQueue) Delete(key int) *QueueItem {
	impl := pq.impl
	if idx, ok := impl.keys[key]; ok {
		// delete
		item := heap.Remove(impl, idx).(QueueItem)
		return &item
	}
	return nil
}

func (pq *PriorityQueue) GetByKey(key int) *QueueItem {
	impl := pq.impl
	if idx, ok := impl.keys[key]; ok {
		item := impl.arr[idx].QueueItem
		return &item
	}
	return nil
}

func (pq *PriorityQueue) UpdatePriority(key int, priority int) bool {
	impl := pq.impl
	if idx, ok := impl.keys[key]; ok {
		impl.arr[idx].Priority = priority
		heap.Fix(impl, idx)
		return true
	}
	return false
}

func (pq *PriorityQueue) UpdatePayload(key int, payload interface{}) bool {
	impl := pq.impl
	if idx, ok := impl.keys[key]; ok {
		impl.arr[idx].Payload = payload
		return true
	}
	return false
}
