package binaryserachtree

import (
	"container/list"
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

func NewBSTree() *BSTree {
	return &BSTree{}
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

func (bst *BSTree) Add(key int64, value interface{}) {
	if bst.root == nil {
		bst.root = &TreeNode{data: &Entry{key, value}}
	} else {
		bst.root.insert(key, value)
	}
	bst.size++
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

func (bst *BSTree) Order(order string) []Entry {
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
