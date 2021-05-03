package main

import (
	"data-structure/src"
	"fmt"
)

func main() {
	stack := src.CreateArrayStack(10)
	for i := 0; i < 10; i++ {
		stack.Push(i+3)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(stack.Pull())
	}
	fmt.Println(stack)
}
