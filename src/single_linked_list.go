package src

import (
	"data-structure/src/node"
	"errors"
	"fmt"
)

type SingleLinkedList struct {
	first   *node.Node
	count   int
	compare func(val1 node.NodeValType, val2 node.NodeValType) bool
}

// 获取一个单链表
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		first: nil,
		count: 0,
	}
}

// 打印链表
func (l *SingleLinkedList) ListPrint() error {
	node, index := l.first, 0
	for node != nil {
		fmt.Println(index, node.GetVal())
		node = node.Next()
		index++
	}
	return nil
}

// 查询index对应的节点
func (l *SingleLinkedList) Get(index int) (*node.Node, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}

	if index == 0 {
		return l.first, nil
	}

	curNode := l.first
	i := 0
	for curNode.Next() != nil && i < index {
		curNode = curNode.Next()
		i++
	}

	if i < index {
		return nil, errors.New("结构异常")
	}

	return curNode, nil
}

// 插入
func (l *SingleLinkedList) Insert(index int, val node.NodeValType) error {
	if index < 0 || index > l.count {
		return errors.New("操作越界")
	}
	n := node.New(val)
	if index == 0 {
		n.SetNext(l.first)
		l.first = n
		l.count++
		return nil
	}
	preNode := l.first
	i := 1
	for preNode.Next() != nil && i < index {
		preNode = preNode.Next()
		i++
	}
	if i < index {
		return errors.New("结构异常")
	}
	n.SetNext(preNode.Next())
	preNode.SetNext(n)
	l.count++
	return nil
}

// 更新
func (l *SingleLinkedList) Update(index int, val node.NodeValType) error {
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
func (l *SingleLinkedList) Delete(index int) (*node.Node, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}
	n := l.first
	if index == 0 {
		l.first = n.Next()
		l.count--
		return n, nil
	}
	preNode := l.first
	i := 1
	for preNode.Next() != nil && i < index {
		preNode = preNode.Next()
		i++
	}
	if i < index {
		return nil, errors.New("结构异常")
	}
	n = preNode.Next()
	preNode.SetNext(n.Next())
	l.count--
	return n, nil
}

// 获取链表长度
func (l *SingleLinkedList) Count() int {
	return l.count
}

// int value 比较
func IntNodeCompare(val1 node.NodeValType, val2 node.NodeValType) bool {
	return val1.(int) > val2.(int)
}
