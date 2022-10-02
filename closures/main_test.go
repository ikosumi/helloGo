package main

import (
	"fmt"
	"testing"
)

func TestFooUsingOneFunc(t *testing.T) {
	wants := 3
	a := Foo()
	a()
	a()
	got := a()

	if wants != got {
		t.Fatal("got", got, "want", wants)
	}
}

func TestFooUsingTwoFunc(t *testing.T) {
	wantsA := 2
	wantsB := 3

	a := Foo()
	b := Foo()

	a()         // => 1
	gotA := a() // => 2

	b()         // => 1
	b()         // => 2
	gotB := b() // => 3

	fmt.Println(gotA)
	fmt.Println(gotB)

	if wantsA != gotA {
		t.Fatalf("gotA %d, wantsA %d, given a() called %d times\n", gotA, wantsA, wantsA)
	}
	if wantsB != gotB {
		t.Fatalf("gotB %d, wantsB %d, given b() called %d times\n", gotB, wantsB, wantsB)
	}
}
