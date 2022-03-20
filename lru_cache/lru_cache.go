package lrucache

import (
	"fmt"
	"rayhaanbhikha/lru-cache/linkedlist"
)

type LRUCache struct {
	list     *linkedlist.LinkList
	values   map[int]*linkedlist.LinkedListNode
	capacity int
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		list:     linkedlist.New(),
		values:   make(map[int]*linkedlist.LinkedListNode),
		capacity: capacity,
	}
}

func (lru *LRUCache) Get(key int) (int, bool) {
	liNode, ok := lru.values[key]
	if !ok {
		return 0, false
	}

	nodeRemoved := liNode.Delete()
	lru.list.AddNode(nodeRemoved)

	return liNode.GetValue(), true
}

func (lru *LRUCache) Set(key int, val int) {
	// TODO: making assumption key and val are the same right now.
	if lru.isFull() {
		nodeDeleted := lru.list.Pop()
		if nodeDeleted == nil {
			return
		}
		delete(lru.values, nodeDeleted.GetValue())
	}
	newNode := lru.list.Add(val)
	lru.values[newNode.GetValue()] = newNode
}

func (lru *LRUCache) String() string {
	str := ""
	for key, node := range lru.values {
		val := node.GetValue()
		str += fmt.Sprintf("%d, %d\n", key, val)
	}
	return str
}

func (lru *LRUCache) isFull() bool {
	return len(lru.values) == lru.capacity
}
