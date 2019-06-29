package minimum_absolute_difference_in_tree

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

func getMinimumDifference(root *TreeNode) int {
	var result []int
	var minint = 65535
	InOrder(&result, root)
	for i := 1; i < len(result); i++ {
		if result[i]-result[i-1] < minint {
			minint = result[i] - result[i-1]
		}
	}
	return minint
}

func InOrder(result *[]int, node *TreeNode) {
	if node != nil {
		InOrder(result, node.Left)
		*result = append(*result, node.Val)
		InOrder(result, node.Right)
	}
}
