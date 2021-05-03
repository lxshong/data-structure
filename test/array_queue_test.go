package test

import (
	"data-structure/src"
	"testing"
)

func TestEnqueue(t *testing.T) {
	l := 10
	stack := src.CreateArrayQueue(l)
	for i := 0; i < l; i++ {
		err := stack.Enqueue(i + 1)
		if err != nil {
			t.Error("enqueue err")
		}
	}
	err := stack.Enqueue(12)
	if err == nil || err.Error() != "full" {
		t.Error("enqueue err")
	}
}

func TestDequeue(t *testing.T) {
	l := 10
	stack := src.CreateArrayQueue(l)
	i := 1
	for ; i < l; i++ {
		err := stack.Enqueue(i)
		if err != nil {
			t.Error("enqueue err")
		}
	}

	for i = 1; i < l; i++ {
		val, err := stack.Dequeue()
		if val.(int) != (i) || err != nil {
			t.Error("dequeue err")
		}
	}
	val, err := stack.Dequeue()
	if val != nil || err == nil || err.Error() != "empty" {
		t.Error("dequeue err")
	}
}
