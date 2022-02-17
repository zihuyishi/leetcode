package main

import "testing"

func TestDivide(t *testing.T) {
	v1 := divide(10, 3)
	if v1 != 3 {
		t.Errorf("expect 3, got %d", v1)
	}
	v2 := divide(-10, 3)
	if v2 != -3 {
		t.Errorf("expect 3, got %d", v2)
	}
	v3 := divide(0, -234)
	if v3 != 0 {
		t.Errorf("expect 0, got %d", v3)
	}

	v4 := divide(-1, 2)
	if v4 != 0 {
		t.Errorf("expect 0, got %d", v4)
	}
}
