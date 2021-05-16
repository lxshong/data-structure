package node

import "errors"

// 双向节点
type DoubleNode struct {
	val  NodeValType
	Next *DoubleNode
	Pre  *DoubleNode
}

func NewDoubleNode(val NodeValType) *DoubleNode {
	return &DoubleNode{
		val: val,
	}
}

// 返回节点值
func (n *DoubleNode) GetVal() NodeValType {
	return n.val
}

// 设置节点值
func (n *DoubleNode) SetVal(val NodeValType) error {
	if n == nil {
		return errors.New("节点为空")
	}
	n.val = val
	return nil
}
