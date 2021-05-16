package test

import (
	"data-structure/src"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := src.NewLinkedList()
	if list.Count() != 0 {
		t.Error("初始化异常1")
	}
	if node, err := list.Get(0); node != nil || err != nil {
		t.Error("初始化异常2")
	}
}

func TestInsert(t *testing.T) {
	list := src.NewLinkedList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	if list.Count() != 4 {
		t.Error("插入异常1")
	}
	node, err := list.Get(0)
	if err != nil || node == nil {
		t.Error("插入异常2")
	}
	if node.GetVal().(int) != 2 {
		t.Error("插入异常3")
	}
	node = node.Next()
	if node.GetVal().(int) != 4 {
		t.Error("插入异常4")
	}
	node = node.Next()
	if node.GetVal().(int) != 3 {
		t.Error("插入异常5")
	}
	node = node.Next()
	if node.GetVal().(int) != 1 {
		t.Error("插入异常6")
	}
	node = node.Next()
	if node != nil {
		t.Error("插入异常7")
	}
}

func TestListPrint(t *testing.T) {
	list := src.NewLinkedList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	list.ListPrint()
}

func TestGet(t *testing.T) {
	list := src.NewLinkedList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	if node, err := list.Get(0); node == nil || node.GetVal().(int) != 2 || err != nil {
		t.Error("查询异常1")
	}
	if node, err := list.Get(1); node == nil || node.GetVal().(int) != 4 || err != nil {
		t.Error("查询异常2")
	}
	if node, err := list.Get(2); node == nil || node.GetVal().(int) != 3 || err != nil {
		t.Error("查询异常3")
	}
	if node, err := list.Get(3); node == nil || node.GetVal().(int) != 1 || err != nil {
		t.Error("查询异常4")
	}
}

func TestUpdate(t *testing.T) {
	list := src.NewLinkedList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	if err := list.Update(0, 12); err != nil {
		t.Error("更新异常1")
	}
	if err := list.Update(1, 14); err != nil {
		t.Error("更新异常2")
	}
	if err := list.Update(2, 13); err != nil {
		t.Error("更新异常3")
	}
	if err := list.Update(3, 11); err != nil {
		t.Error("更新异常4")
	}
	if node, err := list.Get(0); node == nil || node.GetVal().(int) != 12 || err != nil {
		t.Error("更新异常5")
	}
	if node, err := list.Get(1); node == nil || node.GetVal().(int) != 14 || err != nil {
		t.Error("更新异常6")
	}
	if node, err := list.Get(2); node == nil || node.GetVal().(int) != 13 || err != nil {
		t.Error("更新异常7")
	}
	if node, err := list.Get(3); node == nil || node.GetVal().(int) != 11 || err != nil {
		t.Error("更新异常8")
	}
}

func TestDelete(t *testing.T) {
	list := src.NewLinkedList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	if node,err := list.Delete(0); err != nil || node == nil || node.GetVal() != 2 {
		t.Error("删除异常1")
	}
	if node,err := list.Delete(0); err != nil || node == nil || node.GetVal() != 4 {
		t.Error("删除异常2")
	}
	if node,err := list.Delete(0); err != nil || node == nil || node.GetVal() != 3 {
		t.Error("删除异常3")
	}
	if node,err := list.Delete(0); err != nil || node == nil || node.GetVal() != 1 {
		t.Error("删除异常4")
	}
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)
	list.Insert(1, 4)
	if node,err := list.Delete(list.Count()-1); err != nil || node == nil || node.GetVal() != 1 {
		t.Error("删除异常5")
	}
	if node,err := list.Delete(list.Count()-1); err != nil || node == nil || node.GetVal() != 3 {
		t.Error("删除异常6")
	}
	if node,err := list.Delete(list.Count()-1); err != nil || node == nil || node.GetVal() != 4 {
		t.Error("删除异常7")
	}
	if node,err := list.Delete(list.Count()-1); err != nil || node == nil || node.GetVal() != 2 {
		t.Error("删除异常8")
	}
	list.ListPrint()
}