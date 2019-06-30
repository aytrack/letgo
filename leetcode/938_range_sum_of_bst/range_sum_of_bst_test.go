package range_sum_of_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	ast := assert.New(t)
	tree := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 18}}}
	exceptSum := 32
	result := rangeSumBST(tree, 7, 15)
	ast.Equal(exceptSum, result)
}
