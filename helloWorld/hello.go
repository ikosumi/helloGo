package main

import (
	"fmt"
)

const (
	kb = 1 << 10
	mb = 1024 * kb
	gb = 1024 * mb
)

type S string
type arrS []S
type A arrS
type B A

type Student struct {
	name    string
	id      int64
	address string
}

func binary(num int) {
	fmt.Printf("%s\t\t%s\n", "base10", "binary")
	for i := 0; i < 30; i++ {
		fmt.Printf("-")
	}
	fmt.Println("-")
	fmt.Printf("%d\t\t%b\n", num, num)
}

func main() {
	b := 1
	binary(b)
	binary(b >> 2)
	var me B
	fmt.Println(me)
	fmt.Println("a kb is", kb, "bytes")
	fmt.Println("a mb is", mb, "bytes")
	fmt.Println("a gb is", gb, "bytes")
	loops()
}

func loops() {
	ilir := Student{"ilir", 1, "munich"}
	fmt.Println(ilir.name)
	fmt.Printf("| \t%s\t | \t%s\t | \t%s\t |\n", "unicode", "unicode", "unicode")

	for i := 0; i < 200; i++ {
		fmt.Printf("| \t%#U\t | \t%#U\t | \t%#U\t |\n", i, i+10, i+20)
		for i := 0; i < 26*3; i++ {
			fmt.Printf("-")
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("%T\n", `gee`)

}
