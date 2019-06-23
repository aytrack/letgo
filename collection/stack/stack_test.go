package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack(5)

	v1 := s.Pop()
	if v1 != nil {
		t.Errorf("Empty stack should pop nil.")
	}

	v1 = s.Peek()
	if v1 != nil {
		t.Errorf("Empty stack should peek nil.")
	}

	if s.Len() != 0 {
		t.Errorf("Stack len is %v != 0", s.Len())
	}

	if !s.IsEmpty() {
		t.Errorf("Stack should be empty")
	}

	e1 := &Entry{1, "hello"}
	pushed, err := s.Push(e1)
	if pushed && err != nil {
		t.Errorf("Stack should push succeeful.")
	}
	if s.Len() != 1 {
		t.Errorf("Stack len is %v != 1", s.Len())
	}

	v1 = s.Peek()
	if v1.key != 1 {
		t.Errorf("Peek is %v != %v", v1, e1)
	}

	v1 = s.Pop()
	if v1.key != 1 {
		t.Errorf("Pop is %v != %v", v1, e1)
	}

	if s.Len() != 0 {
		t.Errorf("After Pop length is %v != 0", s.Len())
	}

	e2 := &Entry{2, "golang"}
	e3 := &Entry{3, "play"}
	e4 := &Entry{4, "fire"}
	e5 := &Entry{5, "money"}
	e6 := &Entry{6, "double"}

	_, _ = s.Push(e1)
	_, _ = s.Push(e2)
	_, _ = s.Push(e3)
	_, _ = s.Push(e4)
	_, _ = s.Push(e5)
	if s.Len() != 5 {
		t.Errorf("Stack len is %v != 5", s.Len())
	}

	v5 := s.Peek()
	if v5 != e5 {
		t.Errorf("Stack Peek is %v != %v", v5, e5)
	}
	pushed, err = s.Push(e6)
	if pushed && err != nil {
		t.Errorf("Push failed when stack is Full.")
	}

}
