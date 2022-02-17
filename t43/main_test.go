package main

import "testing"

func TestMulti(t *testing.T) {
	s1 := multiply("12", "3")
	if s1 != "36" {
		t.Errorf("expect 36, got %s", s1)
	}
	s2 := multiply("9999", "9898")
	if s2 != "98970102" {
		t.Errorf("expect 98970102, got %s", s2)
	}
	s3 := multiply("234235234532", "0")
	if s3 != "0" {
		t.Errorf("expect 0, got %s", s3)
	}
	s4 := multiply("867923452", "2879847")
	if s4 != "2499486749471844" {
		t.Errorf("expect 2499486749471844, got %s", s4)
	}
}
