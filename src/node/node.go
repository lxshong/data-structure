package node

import "errors"

type NodeValType interface{}

type Node struct {
	val  NodeValType
	next *Node
}

func NewNode(val NodeValType) *Node {
	return &Node{
		val: val,
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

// 获取节点
func (n *Node) Next() *Node {
	return n.next
}

// 设置节点
func (n *Node) SetNext(node *Node) {
	n.next = node
}
