package dijkstra

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Enqueue(QueueItem{Key: 1, Priority: 100})
	pq.Enqueue(QueueItem{Key: 2, Priority: 10})
	pq.Enqueue(QueueItem{Key: 3, Priority: 20})
	pq.Enqueue(QueueItem{Key: 4, Priority: 50})
	pq.Enqueue(QueueItem{Key: 5, Priority: 250})
	pq.Enqueue(QueueItem{Key: 6, Priority: 25})
	pq.Enqueue(QueueItem{Key: 5, Priority: 90}) // update 5 into 90
	for i, ansKey := range []int{2, 3, 6, 4, 5, 1} {
		item := pq.Dequeue()
		if item.Key != ansKey {
			t.Fatalf("ERR: item should be %d, but got: %+v", ansKey, item)
		}
		if pq.Len() != 5-i {
			t.Fatalf("ERR: pq len should be %d, but got: %+v", 5-i, pq.Len())
		}
		t.Logf(" OK: pop item = %+v", item)
	}
}
