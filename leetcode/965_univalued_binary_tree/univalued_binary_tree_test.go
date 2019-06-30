package univalued_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnivauledTree(t *testing.T) {
	ast := assert.New(t)
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}
	tree2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}}

	istrue1 := isUnivalTree(tree1)
	istrue2 := isUnivalTree(tree2)

	ast.True(istrue1)
	ast.False(istrue2)
}
