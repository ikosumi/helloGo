package main

import (
	"fmt"
)

var x, y, z int

type Ilir struct {
	a int
	b int
}

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
}

func main2() {
	ilir := Student{"ilir", 1, "munich"}
	fmt.Println(ilir.name)
	//fmt.Printf("%X", 10223224231)
	//fmt.Printf("%#U", 'ë')
	fmt.Printf("| \t%s\t | \t%s\t | \t%s\t |\n", "unicode", "unicode", "unicode")

	for i := 0; i < 200; i++ {
		fmt.Printf("| \t%#U\t | \t%#U\t | \t%#U\t |\n", i, i+10, i+20)
		for i := 0; i < 26*3; i++ {
			fmt.Printf("-")
		}
		fmt.Println()
	}
	//x = 2
	//x = 3
	//a := 1
	//
	//fmt.Println(x)
	//fmt.Println("Hi Go")
	//
	//fmt.Println(a)

	fmt.Println()
	fmt.Printf("%T\n", `gee`)

}
