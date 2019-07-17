package heap

import (
	"math"
)

type MaxHeap struct {
	Element []int
}

func NewMaxHeap() *MaxHeap {
	element := []int{0: math.MaxInt64}
	return &MaxHeap{Element: element}
}

func (m *MaxHeap) Length() int {
	return len(m.Element) - 1
}

func (m *MaxHeap) Push(data int) {
	m.Element = append(m.Element, data)
	i := m.Length()
	for ; i/2 > 0 && m.Element[i] > m.Element[i/2]; i /= 2 {
		m.Element[i], m.Element[i/2] = m.Element[i/2], m.Element[i]
	}
}

func (m *MaxHeap) Pop() int {
	if m.Length() == 0 {
		return -1
	}
	data := m.Element[1]
	m.Element[1] = m.Element[m.Length()]
	m.Element = m.Element[:m.Length()]
	i := 1
	for {
		maxPos := i
		if i*2 <= m.Length() && m.Element[i] < m.Element[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= m.Length() && m.Element[maxPos] < m.Element[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		m.Element[i], m.Element[maxPos] = m.Element[maxPos], m.Element[i]
		i = maxPos
	}
	return data
}
