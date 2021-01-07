package heaputil

import (
	"testing"
)

func TestGetHeap(t *testing.T) {
    expectedSortedKeys := [3]string{"key1", "key2", "key0"}
	expectedMap := map[string]int{"key1": 10, "key2": 1, "key0": 0}
	dict := map[string]int{"key0": 0, "key1": 10, "key2": 1}

	gotHeap := GetHeap(dict)

	for _, expectedKey := range expectedSortedKeys {
	    expectedVal := expectedMap[expectedKey]
		gotKV := gotHeap.HeapPop()
		if gotKV.key != expectedKey {
			t.Errorf("Wanted key: %v but got: %v", expectedKey, gotKV.key)
		}
		if gotKV.val != expectedVal {
			t.Errorf("Wanted value: %v but got: %v", expectedVal, gotKV.val)
		}
	}
}
