package heaputil

import (
	"container/heap"
	"testing"
)

func TestGetHeap(t *testing.T) {
	expectedMap := map[string]int{"key1": 10, "key2": 1, "key0": 0}
	dict := map[string]int{"key0": 0, "key1": 10, "key2": 1}

	got := getHeap(dict)

	for expectedKey, expectedVal := range expectedMap {
		gotEntry := heap.Pop(got)
		gotKey := gotEntry.(kv).key
		gotVal := gotEntry.(kv).val
		if actualKey != expectedKey{
			t.Errorf("Wanted key: %v but got: %v", gotKey, expectedKey)
		}
		if actualVal != expectedVal{
			t.Errorf("Wanted value: %v but got: %v", gotVal, expectedVal)
		}
	}
}
