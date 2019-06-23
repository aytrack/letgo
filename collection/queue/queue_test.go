package queue

import "testing"

func TestQueue_Len(t *testing.T) {
	q := NewQueue()
	if q.Len() != 0 {
		t.Errorf("Queue Len Failed.\n The len is %d != 0", q.length)
	}

	e1 := &Entry{1, "hello"}
	e2 := &Entry{2, "golang"}
	e3 := &Entry{3, "data"}
	q.Push(e1)
	q.Push(e2)
	q.Push(e3)
	if q.Len() != 3 {
		t.Errorf("Queue Len Failed.\n The len is %d != 3", q.length)
	}
}

func TestQueue_Push(t *testing.T) {
	q := NewQueue()
	e1 := &Entry{1, "hello"}
	e2 := &Entry{2, "golang"}
	e3 := &Entry{3, "data"}
	q.Push(e1)
	q.Push(e2)
	q.Push(e3)
	if q.head.entry != e1 {
		t.Errorf("Queue Push Failed.\n Queue head entry is %v != %v", q.head.entry, e1)
	}
	if q.length != 3 {
		t.Errorf("Queue Push Failed.\n Queue length %v != %v", q.head.entry, e1)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()
	pe1 := q.Pop()
	if pe1 != nil {
		t.Errorf("Queue Pop Failed.\n Queue Entry is %v != %v \n", pe1, nil)
	}

	e1 := &Entry{1, "hello"}
	q.Push(e1)
	pe1 = q.Pop()
	if pe1.key != e1.key && q.head != nil && q.tail != nil {
		t.Errorf("Queue Pop Failed.\n Queue Entry is %v != %v \n", pe1, nil)
	}
	e2 := &Entry{2, "golang"}
	e3 := &Entry{3, "data"}
	q.Push(e1)
	q.Push(e2)
	q.Push(e3)

	pe1 = q.Pop()
	if pe1.key != e1.key {
		t.Errorf("Queue Pop Failed.\n Queue Entry is %v != %v \n", pe1, e1)
	}
	if q.length != 2 {
		t.Errorf("Queue Pop Failed.\n Queue len is %v != %v \n", q.length, 2)
	}
}
func TestQueue_Peek0(t *testing.T) {

	q := NewQueue()
	pe1 := q.Peek()
	if pe1 != nil {
		t.Errorf("Queue Peek Failed.\n Queue Entry is %v != %v \n", pe1, nil)
	}
}

func TestQueue_Peek1(t *testing.T) {

	q := NewQueue()
	e1 := &Entry{1, "hello"}
	q.Push(e1)
	pe1 := q.Peek()
	if pe1.key != e1.key {
		t.Errorf("Queue Peek Failed.\n Queue Entry is %v != %v \n", pe1, e1)
	}
	if q.length != 1 {
		t.Errorf("Queue Peek Failed.\n Queue length is %v != %v \n", q.length, 1)
	}
}
