package eviction

import lib "cache/lib"

type LRUEviction[K comparable] struct {
	doublyLinkedList *lib.DoublyLinkedList[K]

	mapper map[K]*lib.Node[K]
}

func New[K comparable]() *LRUEviction[K] {

	return &LRUEviction[K]{
		doublyLinkedList: lib.NewtDoublyList[K](),
		mapper:           make(map[K]*lib.Node[K]),
	}
}

func (lru *LRUEviction[K]) KeyAccessed(key K) {

	if node, ok := lru.mapper[key]; ok {
		lru.doublyLinkedList.DetachNode(node)
		lru.doublyLinkedList.AddEndNodeDLL(node)
	} else {

		newNode := lib.NewNode(key)
		lru.doublyLinkedList.AddEndNodeDLL(newNode)
		lru.mapper[key] = newNode
	}
}

func (lru *LRUEviction[K]) Evict() K {

	firstNode := lru.doublyLinkedList.Head
	lru.doublyLinkedList.DetachNode(firstNode)
	return firstNode.GetKey()

}
