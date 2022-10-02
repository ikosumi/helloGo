package main

import (
	"testing"
)

/*
func sum(<slice>) int {...}
func sumFn(<fn>, <slice>) int {...}

e.g.
func sum([]int{1, 2, 3, 4})
=> 10

func sumFn(even, []int{1, 2, 3, 4})
=> 6

func sumFn(odd, []int{1, 2, 3, 4})
=> 4
*/

func TestSum(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	wants := 10
	got := Sum(inputs...)

	if wants != got {
		t.Fatalf("got %d, want %d, given sum(%v)\n", got, wants, inputs)
	}
}
