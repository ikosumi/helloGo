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
		t.Fatalf("got %d, want %d, given Sum(%v)\n", got, wants, inputs)
	}
}

func TestSumFnWithEven(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	wants := 6
	got := SumFn(Even, inputs...)

	if wants != got {
		t.Fatalf("got %d, want %d, given SumFn(Even, %v)\n", got, wants, inputs)
	}
}

func TestSumFnWithOdd(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	wants := 4
	got := SumFn(Odd, inputs...)

	if wants != got {
		t.Fatalf("got %d, want %d, given SumFn(Odd, %v)\n", got, wants, inputs)
	}
}
