package invert_binary_tree

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用递归方法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 使用迭代方法
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		elem := queue.Front()
		queue.Remove(elem)
		p := elem.Value.(*TreeNode)
		p.Left, p.Right = p.Right, p.Left
		if p.Left != nil {
			queue.PushBack(p.Left)
		}
		if p.Right != nil {
			queue.PushBack(p.Right)
		}
	}
	return root
}
