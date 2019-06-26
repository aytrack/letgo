package binaryserachtree

import (
	"container/list"
	"errors"
)

type BSTree struct {
	root *TreeNode
	size int
}

type TreeNode struct {
	data        *Entry
	left, right *TreeNode
}

type Entry struct {
	key   int64
	value interface{}
}

type MyQueue struct {
	List *list.List
}

type MyStack struct {
	List *list.List
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func (s *MyStack) Push(v interface{}) {
	s.List.PushBack(v)
}

func (s *MyStack) Pop() interface{} {
	if elem := s.List.Back(); elem != nil {
		s.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (q *MyQueue) Push(v interface{}) {
	q.List.PushBack(v)
}

func (q *MyQueue) Pop() interface{} {
	if elem := q.List.Front(); elem != nil {
		q.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (bst *BSTree) InsertRecur(key int64, value interface{}) {
	if bst.root == nil {
		bst.root = &TreeNode{data: &Entry{key, value}}
	} else {
		bst.root.insert(key, value)
	}
	bst.size++
}

func (bst *BSTree) Insert(key int64, value interface{}) {
	newNode := &TreeNode{data: &Entry{key: key, value: value}}
	if bst.root == nil {
		bst.root = newNode
		return
	}
	node := bst.root
	for node != nil {
		if key > node.data.key {
			if node.right == nil {
				node.right = newNode
				return
			}
			node = node.right
		}
		if key < node.data.key {
			if node.left == nil {
				node.left = newNode
				return
			}
			node = node.left
		}
	}
}

func (bst *BSTree) Find(key int64) *Entry {
	node := bst.root
	for {
		if node == nil {
			return nil
		} else if node.data.key == key {
			return node.data
		} else if node.data.key < key {
			node = node.right
		} else if node.data.key > key {
			node = node.left
		}
	}
}

func (bst *BSTree) OrderRecur(order string) []Entry {
	var result []Entry
	switch order {
	case "pre":
		bst.root.preOrder(&result)
	case "in":
		bst.root.inOrder(&result)
	case "post":
		bst.root.postOrder(&result)
	default:
		bst.root.preOrder(&result)
	}
	return result
}

func (bst *BSTree) PreOrder() []Entry {
	var result []Entry
	s := &MyStack{List: list.New()}
	node := bst.root

	for node != nil || s.List.Len() != 0 {
		for node != nil {
			result = append(result, *node.data)
			s.Push(node)
			node = node.left
		}
		if s.List.Len() != 0 {
			node = s.Pop().(*TreeNode)
			node = node.right
		}
	}
	return result
}

func (bst *BSTree) InOrder() []Entry {
	var result []Entry
	s := &MyStack{list.New()}
	node := bst.root
	for node != nil || s.List.Len() != 0 {
		for node != nil {
			s.Push(node)
			node = node.left
		}
		if s.List.Len() != 0 {
			node = s.Pop().(*TreeNode)
			result = append(result, *node.data)
			node = node.right
		}
	}
	return result
}

func (bst *BSTree) PostOrder() []Entry {
	s := &MyStack{List: list.New()}
	node := bst.root
	var prePop *TreeNode
	var result []Entry
	for (node != nil || s.List.Len() != 0) && prePop != bst.root {
		for node != nil {
			s.Push(node)
			node = node.left
		}
		for s.List.Len() != 0 {
			node = s.Pop().(*TreeNode)
			if node.right == nil || node.right == prePop {
				result = append(result, *node.data)
				prePop = node
			} else {
				s.Push(node)
				node = node.right
				break
			}
		}
	}
	return result
}
func (bst *BSTree) LevelOrder() []Entry {
	queue := MyQueue{list.New()}
	var result []Entry
	node := bst.root
	queue.Push(node)
	for node != nil && queue.List.Len() != 0 {
		node = queue.Pop().(*TreeNode)
		result = append(result, *node.data)
		if node.left != nil {
			queue.Push(node.left)
		}
		if node.right != nil {
			queue.Push(node.right)
		}
	}
	return result
}

func (bst *BSTree) Min() *Entry {
	node := bst.root
	for node.left != nil {
		node = node.left
	}
	return node.data
}

func (bst *BSTree) Max() *Entry {
	node := bst.root
	for node.right != nil {
		node = node.right
	}
	return node.data
}

func (bst *BSTree) Delete(data *Entry) (bool, error) {
	node := bst.root
	var parentNode *TreeNode
	for node != nil && node.data.key != data.key {
		if node.data.key > data.key {
			parentNode = node
			node = node.left
		}
		if node.data.key < data.key {
			parentNode = node
			node = node.right
		}
	}
	if node == nil {
		return false, errors.New("data not exists.")
	}
	if node.right != nil && node.left != nil {
		minNode := node.right
		minNodeParent := node
		for minNode.left != nil {
			minNode = minNode.left
			minNodeParent = minNode
		}
		node.data = minNode.data
		node = minNode
		parentNode = minNodeParent
	}
	var child *TreeNode
	if node.left != nil {
		child = node.left
	} else if node.right != nil {
		child = node.right
	} else {
		child = nil
	}

	if parentNode == nil {
		bst.root = nil
		return true, nil
	} else if parentNode.left == node {
		parentNode.left = child
		return true, nil
	} else if parentNode.right == node {
		parentNode.right = child
		return true, nil
	}
	return false, errors.New("Delete node failed.")
}

func (n *TreeNode) preOrder(result *[]Entry) {
	if n != nil {
		*result = append(*result, *n.data)
		n.left.preOrder(result)
		n.right.preOrder(result)
	}
}

func (n *TreeNode) inOrder(result *[]Entry) {
	if n != nil {
		n.left.inOrder(result)
		*result = append(*result, *n.data)
		n.right.inOrder(result)
	}
}

func (n *TreeNode) postOrder(result *[]Entry) {
	if n != nil {
		n.left.postOrder(result)
		n.right.postOrder(result)
		*result = append(*result, *n.data)
	}
}

func (n *TreeNode) insert(key int64, value interface{}) {
	node := &TreeNode{data: &Entry{key, value}}
	if key > n.data.key {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insert(key, value)
		}
	}
	if key < n.data.key {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insert(key, value)
		}
	}
}
