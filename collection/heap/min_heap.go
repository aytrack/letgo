package heap

import "math"

type MinHeap struct {
	Element []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Element: []int{0: math.MinInt64}}
}

func (m *MinHeap) Length() int {
	return len(m.Element) - 1
}

func (m *MinHeap) Push(data int) {
	m.Element = append(m.Element, data)
	for i := m.Length(); m.Element[i] < m.Element[i/2] && i/2 > 0; i /= 2 {
		m.Element[i], m.Element[i/2] = m.Element[i/2], m.Element[i]
	}
}

func (m *MinHeap) Pop() int {
	if m.Length() == 0 {
		return -1
	}
	data := m.Element[1]
	m.Element[1] = m.Element[m.Length()]
	m.Element = m.Element[:m.Length()]
	i := 1
	for {
		minPos := i
		if i*2 <= m.Length() && m.Element[i] > m.Element[i*2] {
			minPos = i * 2
		}
		if i*2+1 <= m.Length() && m.Element[minPos] > m.Element[i*2+1] {
			minPos = i*2 + 1
		}
		if minPos == i {
			break
		}
		m.Element[i], m.Element[minPos] = m.Element[minPos], m.Element[i]
		i = minPos
	}
	return data
}
