package invert_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	ast := assert.New(t)
	var result []int
	tree := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	invertTree := invertTree(tree)
	InOrder(&result, invertTree)
	exceptResult := []int{9, 7, 6, 4, 3, 2, 1}
	ast.Equal(result, exceptResult)

	result = []int{}
	InvertTree := InvertTree(tree)
	InOrder(&result, InvertTree)
	exceptResult = []int{1, 2, 3, 4, 6, 7, 9}
	ast.Equal(result, exceptResult)
}

func InOrder(result *[]int, node *TreeNode) {
	if node != nil {
		InOrder(result, node.Left)
		*result = append(*result, node.Val)
		InOrder(result, node.Right)
	}
}
