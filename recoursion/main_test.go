package main

import "testing"

func TestFactorialOf5(t *testing.T) {
	wants := 120
	got := fact(5)

	if wants != got {
		t.Fatal("got", got, "wants", wants)
	}
}
