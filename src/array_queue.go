package src

import "errors"

// 顺序栈
type arrayQueue struct {
	nodes []interface{}
	len   int
	count   int
	head  int
	tail  int
}

// 创建栈
func CreateArrayQueue(n int) *arrayQueue {
	return &arrayQueue{
		nodes: make([]interface{}, n),
		len:   n,
		count:   0,
		head:  0,
		tail:  0,
	}
}

// 入队列
func (receiver *arrayQueue) Enqueue(value interface{}) error {
	if receiver.len == receiver.count {
		return errors.New("full")
	}
	receiver.nodes[receiver.tail] = value
	receiver.tail++
	receiver.count++
	if receiver.tail >= receiver.len {
		receiver.tail = 0
	}
	return nil
}

// 出队列
func (receiver *arrayQueue) Dequeue() (interface{}, error) {
	if receiver.count == 0 {
		return nil, errors.New("empty")
	}
	value := receiver.nodes[receiver.head]
	receiver.head++
	receiver.count--
	if receiver.head >= receiver.len {
		receiver.head = 0
	}
	return value, nil
}
