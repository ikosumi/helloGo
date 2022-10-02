package main

import "fmt"

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	var errs []bool

	printSlice(errs)
	println(errs)

	for i := 0; i < len(m); i++ {
		println(m[i])
	}

	println(m[1], m[2])

	if val, ok := m[1]; ok {
		println(val, ok)
	}

	if name := "ilir"; true {
		println(name)
	}

	var name string

	if name == "ilir" {
		println(name)
	}

}
func printSlice(s []bool) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
