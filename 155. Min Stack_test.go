package main

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin()) // return -3
	minStack.Pop()
	println(minStack.Top())    // return 0
	println(minStack.GetMin()) // return -2
}
