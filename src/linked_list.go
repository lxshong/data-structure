package src

import (
	"errors"
	"fmt"
)

type NodeValType interface{}

type LinkedList struct {
	first   *Node
	count   int
	compare func(val1 NodeValType, val2 NodeValType) bool
}

type Node struct {
	val  NodeValType
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		first: nil,
		count: 0,
	}
}

// 返回节点值
func (n *Node) GetVal() NodeValType {
	return n.val
}

// 设置节点值
func (n *Node) SetVal(val NodeValType) error {
	if n == nil {
		return errors.New("节点为空")
	}
	n.val = val
	return nil
}

// 下一个节点
func (n *Node) Next() *Node {
	return n.next
}

// 打印链表
func (l *LinkedList) ListPrint() error {
	node, index := l.first, 0
	for node != nil {
		fmt.Println(index, node.val)
		node = node.next
		index++
	}
	return nil
}

// 查询index对应的节点
func (l *LinkedList) Get(index int) (*Node, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}

	if index == 0 {
		return l.first, nil
	}

	curNode := l.first
	i := 0
	for curNode.next != nil && i < index {
		curNode = curNode.next
		i++
	}

	if i < index {
		return nil, errors.New("结构异常")
	}

	return curNode, nil
}

// 插入
func (l *LinkedList) Insert(index int, val NodeValType) error {
	if index < 0 || index > l.count {
		return errors.New("操作越界")
	}
	node := &Node{
		val: val,
	}
	if index == 0 {
		node.next = l.first
		l.first = node
		l.count++
		return nil
	}
	preNode := l.first
	i := 1
	for preNode.next != nil && i < index {
		preNode = preNode.next
		i++
	}
	if i < index {
		return errors.New("结构异常")
	}
	node.next, preNode.next = preNode.next, node
	l.count++
	return nil
}

// 更新
func (l *LinkedList) Update(index int, val NodeValType) error {
	if node, err := l.Get(index); err != nil || node == nil {
		return err
	} else {
		if err := node.SetVal(val); err != nil {
			return err
		}
	}
	return nil
}

// 删除
func (l *LinkedList) Delete(index int) (*Node, error) {
	if index < 0 || index > l.count-1 {
		return nil, errors.New("操作越界")
	}
	node := l.first
	if index == 0 {
		l.first = node.next
		l.count--
		return node, nil
	}
	preNode := l.first
	i := 1
	for preNode.next != nil && i < index {
		preNode = preNode.next
		i++
	}
	if i < index {
		return nil, errors.New("结构异常")
	}
	node = preNode.next
	preNode.next = node.next
	l.count--
	return node, nil
}

// 获取链表长度
func (l *LinkedList) Count() int {
	return l.count
}

// int value 比较
func IntNodeCompare(val1 NodeValType, val2 NodeValType) bool {
	return val1.(int) > val2.(int)
}
