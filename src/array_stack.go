package src

import "errors"

// 顺序栈
type arrayStack struct {
	nodes []interface{}
	len   int
	count int
}

// 创建栈
func CreateArrayStack(n int) *arrayStack {
	return &arrayStack{
		nodes: make([]interface{}, n),
		len:   n,
		count: 0,
	}
}

// 出栈
func (receiver *arrayStack) Push(value interface{}) error {
	if receiver.count == receiver.len {
		return errors.New("full")
	}
	receiver.nodes[receiver.count] = value
	receiver.count++
	return nil
}

// 入栈
func (receiver *arrayStack) Pull() (interface{}, error) {
	if receiver.count == 0 {
		return nil, errors.New("empty")
	}
	receiver.count--
	value := receiver.nodes[receiver.count]
	return value, nil
}
