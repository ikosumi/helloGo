package main

import "fmt"

func main() {
	bar := Foo()
	buzz := Foo()

	fmt.Println(bar())  // => 1
	fmt.Println(bar())  // => 2
	fmt.Println(buzz()) // => 1
	fmt.Println(buzz()) // => 2
	fmt.Println(bar())  // => 3
}

func Foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
