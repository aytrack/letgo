package stack

import "errors"

type Stack struct {
	length int
	top    *Node
	count  int
}

type Node struct {
	value *Entry
	prev  *Node
}

type Entry struct {
	key   int64
	value interface{}
}

func NewStack(count int) *Stack {
	return &Stack{length: 0, top: nil, count: count}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) IsFull() bool {
	return s.length == s.count
}

func (s *Stack) Push(entry *Entry) (bool, error) {
	if s.IsFull() {
		return false, errors.New("Stack is Full.")
	}
	node := &Node{value: entry, prev: s.top}
	s.top = node
	s.length++
	return true, nil
}

func (s *Stack) Pop() *Entry {
	if s.IsEmpty() {
		return nil
	}
	node := s.top
	s.top = node.prev
	s.length--
	return node.value
}

func (s *Stack) Peek() *Entry {
	if s.IsEmpty() {
		return nil
	}
	return s.top.value
}
