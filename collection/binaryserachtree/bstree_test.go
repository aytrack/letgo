package binaryserachtree

import (
	"testing"
)

func TestBSTree_PreOrder(t *testing.T) {
	bst := NewBSTree()

	v0 := bst.Find(1)
	if v0 != nil {
		t.Errorf("Empty tree find Failed.")
	}

	v1 := bst.LevelOrder()
	if v1 != nil {
		t.Errorf("Empty tree level travel failed.")
	}

	bst.Add(3, "a")
	bst.Add(2, "b")
	bst.Add(5, "c")
	bst.Add(1, "d")
	bst.Add(4, "e")

	preExcept := []Entry{{3, "a"}, {2, "b"}, {1, "d"}, {5, "c"}, {4, "e"}}
	inExcept := []Entry{{1, "d"}, {2, "b"}, {3, "a"}, {4, "e"}, {5, "c"}}
	postExcept := []Entry{{1, "d"}, {2, "b"}, {4, "e"}, {5, "c"}, {3, "a"}}
	levelExcept := []Entry{{3, "a"}, {2, "b"}, {5, "c"}, {1, "d"}, {4, "e"}}

	preOrderv := bst.Order("s")
	if !compareList(preExcept, preOrderv) {
		t.Errorf("Default order is %v != %v", preOrderv, preExcept)
	}

	preOrderv = bst.Order("pre")
	if !compareList(preExcept, preOrderv) {
		t.Errorf("PreOrder is %v != %v", preOrderv, preExcept)
	}

	inOrderv := bst.Order("in")
	if !compareList(inExcept, inOrderv) {
		t.Errorf("InOrder is %v != %v", inOrderv, inExcept)
	}

	levelOrderv := bst.LevelOrder()
	if !compareList(levelExcept, levelOrderv) {
		t.Errorf("LevelOrder is %v != %v", levelOrderv, levelExcept)
	}

	postOrderv := bst.Order("post")
	if !compareList(postExcept, postOrderv) {
		t.Errorf("PostOrder is %v != %v", postOrderv, postExcept)
	}

	v2 := bst.Find(2)
	if v2.key != 2 {
		t.Errorf("Find v1 key is %v != 2", v2.key)
	}

	v3 := bst.Find(4)
	if v3.key != 4 {
		t.Errorf("Find v1 key is %v != 4", v3.key)
	}

	min := bst.Min()
	if min.key != 1 {
		t.Errorf("Min's key is %d != 1", min)
	}

	max := bst.Max()
	if max.key != 5 {
		t.Errorf("Max's key is %d != 5", max)
	}

}

func compareList(l1, l2 []Entry) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
