package src

import (
	"data-structure/src/node"
	"errors"
	"fmt"
)

// 双向链表
type DoubleLinkedList struct {
	first   *node.DoubleNode
	end     *node.DoubleNode
	count   int
	compare func(val1 node.NodeValType, val2 node.NodeValType) bool
}

// 获取一个双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		first: nil,
		count: 0,
	}
}

// 打印链表
func (l *DoubleLinkedList) ListPrint() error {
	node, index := l.first, 0
	for node != nil {
		fmt.Println(index, node.GetVal())
		node = node.Next
		index++
	}
	return nil
}

// 查询index对应的节点
func (l *DoubleLinkedList) Get(index int) (*node.DoubleNode, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}

	curNode := l.first
	i := 0
	for curNode.Next != nil && i < index {
		curNode = curNode.Next
		i++
	}

	if i < index {
		return nil, errors.New("结构异常")
	}

	return curNode, nil
}

// 插入
func (l *DoubleLinkedList) Insert(index int, val node.NodeValType) error {
	if index < 0 || index > l.count {
		return errors.New("操作越界")
	}
	n := node.NewDoubleNode(val)
	if index == 0 {
		if l.first != nil {
			n.Next = l.first
			l.first.Pre = n
		}
		l.first = n
		l.count++
		return nil
	}
	curNode := l.first
	i := 0
	for curNode.Next != nil && i < index {
		curNode = curNode.Next
		i++
	}
	if i < index {
		return errors.New("结构异常")
	}
	n.Next = curNode
	n.Pre = curNode.Pre
	curNode.Pre.Next = n
	curNode.Pre = n
	l.count++
	return nil
}

// 更新
func (l *DoubleLinkedList) Update(index int, val node.NodeValType) error {
	if n, err := l.Get(index); err != nil || n == nil {
		return err
	} else {
		if err := n.SetVal(val); err != nil {
			return err
		}
	}
	return nil
}

// 删除
func (l *DoubleLinkedList) Delete(index int) (*node.DoubleNode, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}
	n := l.first
	if index == 0 {
		l.first = n.Next
		l.count--
		if n.Next != nil {
			n.Next.Pre = nil
		}
		n.Next = nil
		return n, nil
	}
	i := 0
	for n.Next != nil && i < index {
		n = n.Next
		i++
	}
	if i < index {
		return nil, errors.New("结构异常")
	}
	n.Pre.Next = n.Next
	if n.Next != nil {
		n.Next.Pre = n.Pre
	}
	l.count--
	return n, nil
}

// 获取链表长度
func (l *DoubleLinkedList) Count() int {
	return l.count
}
