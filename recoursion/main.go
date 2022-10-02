package main

import "fmt"

func main() {
	fmt.Println(fact(5))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
