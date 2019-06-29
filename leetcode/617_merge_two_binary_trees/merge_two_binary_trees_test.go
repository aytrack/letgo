package merge_two_binary_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	ast := assert.New(t)
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	exceptPreOrder := []int{3, 4, 5, 4, 5, 7}
	var result []int
	PreOrder(&result, mergeTrees(t1, t2))
	ast.Equal(result, exceptPreOrder)
}

func PreOrder(result *[]int, node *TreeNode) {
	if node != nil {
		*result = append(*result, node.Val)
		PreOrder(result, node.Left)
		PreOrder(result, node.Right)
	}
}
