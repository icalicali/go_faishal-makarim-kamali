package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 12
	var actual int = add(10, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

func TestSubstract(t *testing.T) {
	var expected int = 8
	var actual int = add(10, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	var expected int = 20
	var actual int = add(10, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

func TestDivison(t *testing.T) {
	var expected int = 5
	var actual int = add(10, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}
