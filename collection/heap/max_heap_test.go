package heap

import "testing"

func TestMaxHeap_Push(t *testing.T) {
	maxHeap := NewMaxHeap()

	maxHeap.Push(10)
	maxHeap.Push(16)
	maxHeap.Push(7)
	maxHeap.Push(3)
	maxHeap.Push(33)
	maxHeap.Push(8)

	t.Log(maxHeap)

	for i := 0; i <= 6; i++ {
		t.Log(maxHeap.Pop())
	}
}
