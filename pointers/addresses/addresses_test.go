package main

import "testing"

func TestUpdateValueInAddr(t *testing.T) {
	input := 11
	wants := 42
	UpdateValueInAddr(&input, wants)
	got := input
	if wants != got {
		t.Fatalf("got %d, want %d\n", got, wants)
	}
}
