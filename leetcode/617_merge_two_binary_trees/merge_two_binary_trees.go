package merge_two_binary_trees

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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	if t1 == nil {
		root = t2
	}
	if t2 == nil {
		root = t1
	}
	if t1 != nil && t2 != nil {
		root.Val = t1.Val + t2.Val
		root.Left = mergeTrees(t1.Left, t2.Left)
		root.Right = mergeTrees(t1.Right, t2.Right)
	}
	return root
}
