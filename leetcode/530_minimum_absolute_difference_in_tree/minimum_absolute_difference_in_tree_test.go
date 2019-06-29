package minimum_absolute_difference_in_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	ast := assert.New(t)
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}
	result := getMinimumDifference(tree)
	ast.Equal(result, 1)
}
