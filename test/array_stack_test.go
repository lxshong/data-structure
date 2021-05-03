package test

import (
	"data-structure/src"
	"testing"
)

func TestPush(t *testing.T) {
	l := 10
	stack := src.CreateArrayStack(l)
	for i := 0; i < l; i++ {
		err := stack.Push(i + 1)
		if err != nil {
			t.Error("push err")
		}
	}
	err := stack.Push(12)
	if err == nil || err.Error() != "full" {
		t.Error("push err")
	}
}

func TestPull(t *testing.T) {
	l := 10
	stack := src.CreateArrayStack(l)
	i := 1
	for ; i < l; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Error("pull err")
		}
	}

	for i--; i > 0; i-- {
		val, err := stack.Pull()
		if val.(int) != (i) || err != nil {
			t.Error("pull err")
		}
	}
	val, err := stack.Pull()
	if val != nil || err == nil || err.Error() != "empty" {
		t.Error("pull err")
	}
}
