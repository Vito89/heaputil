// Package heaputil contains utility functions for working with structures and container/heap
package heaputil

import (
	"container/heap"
)

type KeyValue struct {
	key string
	val int
}

type KVHeap []KeyValue

func (h KVHeap) Len() int           { return len(h) }
func (h KVHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h KVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(KeyValue))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// getHeap returns heap of map
func GetHeap(m map[string]int) *KVHeap {
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, KeyValue{k, v})
	}
	return h
}