package test

import (
	"data-structure/src"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := src.InsertSort([]interface{}{5, 4, 2, 3, 1})
	for i := 0; i < 5; i++ {
		if arr[i].(int) != i+1 {
			t.Error("insert sort error")
		}
	}
}
