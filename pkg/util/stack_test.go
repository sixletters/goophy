package util

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}

	// Test Push() and Size() methods
	s.Push(1)
	s.Push("two")
	s.Push(3.0)

	if s.Size() != 3 {
		t.Errorf("Expected stack size to be 3, but got %d", s.Size())
	}

	// Test Pop() method
	if s.Pop() != 3.0 {
		t.Errorf("Expected popped item to be 3.0, but got %v", s.Pop())
	}
	if s.Pop() != "two" {
		t.Errorf("Expected popped item to be 'two', but got %v", s.Pop())
	}
	if s.Pop() != 1 {
		t.Errorf("Expected popped item to be 1, but got %v", s.Pop())
	}
	if s.Pop() != nil {
		t.Errorf("Expected popped item to be nil, but got %v", s.Pop())
	}

	// Test Peek() method
	s.Push("hello")
	s.Push("world")

	if s.Peek() != "world" {
		t.Errorf("Expected top item to be 'world', but got %v", s.Peek())
	}
	s.Pop()
	if s.Peek() != "hello" {
		t.Errorf("Expected top item to be 'hello', but got %v", s.Peek())
	}
}

func TestStackWithCustomType(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	s := Stack{}

	// Push some items onto the stack
	s.Push(person{"Alice", 30})
	s.Push(person{"Bob", 25})
	s.Push(person{"Charlie", 40})

	// Test Peek() function
	if p, ok := s.Peek().(person); ok {
		if p.name != "Charlie" || p.age != 40 {
			t.Errorf("Peek() = %v, want %v", p, person{"Charlie", 40})
		}
	} else {
		t.Errorf("Peek() returned unexpected type")
	}

	// Test Pop() function
	if p, ok := s.Pop().(person); ok {
		if p.name != "Charlie" || p.age != 40 {
			t.Errorf("Pop() = %v, want %v", p, person{"Charlie", 40})
		}
	} else {
		t.Errorf("Pop() returned unexpected type")
	}

	// Test Size() function
	if size := s.Size(); size != 2 {
		t.Errorf("Size() = %d, want %d", size, 2)
	}
}

func TestStack_Pop_Empty(t *testing.T) {
	s := Stack{}
	v := s.Pop()
	if v != nil {
		t.Errorf("Expected nil, got %v", v)
	}
}
