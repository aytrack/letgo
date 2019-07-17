package heap

import "testing"

func TestMinHeap_Push(t *testing.T) {
	minHeap := NewMinHeap()

	minHeap.Push(10)
	minHeap.Push(16)
	minHeap.Push(7)
	minHeap.Push(3)
	minHeap.Push(33)
	minHeap.Push(8)

	t.Log(minHeap)

	for i := 0; i <= 6; i++ {
		t.Log(minHeap.Pop())
	}
}
