package chess

import "testing"

func TestStack(t *testing.T) {
	uut := new(Stack)
	expected := 23
	uut.Push(expected)

	len := uut.Len()
	if len != 1 {
		t.Errorf("Expected 1 but got %d", len)
	}
	actual := uut.Pop()
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	len = uut.Len()
	if len != 0 {
		t.Errorf("Expected 0 but got %d", len)
	}

	actual = uut.Pop()
	if actual != nil {
		t.Errorf("Expected nil but got %v", actual)
	}
}
