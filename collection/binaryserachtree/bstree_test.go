package binaryserachtree

import (
	"testing"
)

func TestBSTree(t *testing.T) {
	bst := NewBSTree()

	v0 := bst.Find(1)
	if v0 != nil {
		t.Errorf("Empty tree find Failed.")
	}

	v1 := bst.LevelOrder()
	if v1 != nil {
		t.Errorf("Empty tree level travel failed.")
	}

	bst.InsertRecur(3, "a")
	bst.InsertRecur(2, "b")
	bst.InsertRecur(5, "c")
	bst.InsertRecur(1, "d")
	bst.InsertRecur(4, "e")

	preExcept := []Entry{{3, "a"}, {2, "b"}, {1, "d"}, {5, "c"}, {4, "e"}}
	inExcept := []Entry{{1, "d"}, {2, "b"}, {3, "a"}, {4, "e"}, {5, "c"}}
	postExcept := []Entry{{1, "d"}, {2, "b"}, {4, "e"}, {5, "c"}, {3, "a"}}
	levelExcept := []Entry{{3, "a"}, {2, "b"}, {5, "c"}, {1, "d"}, {4, "e"}}

	preOrderv := bst.OrderRecur("s")
	if !compareList(preExcept, preOrderv) {
		t.Errorf("Default order is %v != %v", preOrderv, preExcept)
	}

	preOrderv = bst.OrderRecur("pre")
	if !compareList(preExcept, preOrderv) {
		t.Errorf("PreOrderRecur is %v != %v", preOrderv, preExcept)
	}

	preOrderv = bst.PreOrder()
	if !compareList(preExcept, preOrderv) {
		t.Errorf("PreOrder is %v != %v", preOrderv, preExcept)
	}

	inOrderv := bst.OrderRecur("in")
	if !compareList(inExcept, inOrderv) {
		t.Errorf("InOrderRecur is %v != %v", inOrderv, inExcept)
	}

	inOrderv = bst.InOrder()
	if !compareList(inExcept, inOrderv) {
		t.Errorf("InOrder is %v != %v", inOrderv, inExcept)
	}

	levelOrderv := bst.LevelOrder()
	if !compareList(levelExcept, levelOrderv) {
		t.Errorf("LevelOrder is %v != %v", levelOrderv, levelExcept)
	}

	postOrderv := bst.OrderRecur("post")
	if !compareList(postExcept, postOrderv) {
		t.Errorf("PostOrder is %v != %v", postOrderv, postExcept)
	}

	postOrderv = bst.PostOrder()
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

	ok, _ := bst.Delete(&Entry{5, "c"})
	inExcept = []Entry{{1, "d"}, {2, "b"}, {3, "a"}, {4, "e"}}
	afterDelete := bst.InOrder()
	if !ok || !compareList(inExcept, afterDelete) {
		t.Errorf("After delete is %v != %v", afterDelete, inExcept)
	}

	bst.Insert(5, "c")
	ok, _ = bst.Delete(&Entry{1, "d"})
	inExcept = []Entry{{2, "b"}, {3, "a"}, {4, "e"}, {5, "c"}}
	afterDelete = bst.InOrder()
	if !ok || !compareList(inExcept, afterDelete) {
		t.Errorf("After delete is %v != %v", afterDelete, inExcept)
	}

	bst.Insert(1, "d")
	ok, _ = bst.Delete(&Entry{3, "a"})
	inExcept = []Entry{{1, "d"}, {2, "b"}, {4, "e"}, {5, "c"}}
	afterDelete = bst.InOrder()
	if !ok || !compareList(inExcept, afterDelete) {
		t.Errorf("After delete is %v != %v", afterDelete, inExcept)
	}

	bst1 := NewBSTree()
	bst1.Insert(10, "aa")
	if bst1.root.data.key != 10 {
		t.Errorf("bst1 root key is %d != 10", bst1.root.data.key)
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
