package eviction

import (
	lib "cache/lib"
	"container/heap"
)

type LFUEviction[K comparable] struct {
	pq *lib.PriorityQueue[K]

	mapper map[K]*lib.Item[K]
}

func NewLFUPolicy[K comparable]() *LFUEviction[K] {

	return &LFUEviction[K]{
		pq:     &lib.PriorityQueue[K]{},
		mapper: make(map[K]*lib.Item[K]),
	}
}

func (lfu *LFUEviction[K]) KeyAccessed(key K) {

	if item, ok := lfu.mapper[key]; ok {
		lfu.pq.Update(item, key, item.GetPriority()+1)
	} else {
		newItem := lib.NewItem(key)
		heap.Push(lfu.pq, newItem)
		lfu.mapper[key] = newItem
	}
}

func (lfu *LFUEviction[K]) Evict() K {

	lfuItem := (heap.Pop(lfu.pq)).(*lib.Item[K])

	key := lfuItem.GetKey()
	delete(lfu.mapper, lfuItem.GetKey())

	return key

}
