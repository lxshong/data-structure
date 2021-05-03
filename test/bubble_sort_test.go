package test

import (
	"data-structure/src"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := src.BubbleSort([]interface{}{5, 4, 2, 3, 1})
	fmt.Println(arr)
	for i := 0; i < 5; i++ {
		if arr[i].(int) != i+1 {
			t.Error("bubble sort error")
		}
	}
}
