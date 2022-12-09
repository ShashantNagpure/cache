package lib

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item[K any] struct {
	value    K   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem[K any](value K) *Item[K] {

	return &Item[K]{value: value, priority: 1}
}
func (item *Item[K]) GetPriority() int {
	return item.priority
}
func (item *Item[K]) GetKey() K {
	return item.value
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[K any] []*Item[K]

func NewHeap[K any]() PriorityQueue[K] {

	return PriorityQueue[K]{}
}

func (pq PriorityQueue[K]) Len() int { return len(pq) }

func (pq PriorityQueue[K]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[K]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[K]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[K])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[K]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[K]) Update(item *Item[K], value K, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item[string]{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item[string]{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
