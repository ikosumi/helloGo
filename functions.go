package main

import "fmt"

func main() {
	inputs := []int{1, 2, 3, 4}
	fmt.Println("the inputs slice is", inputs)
	fmt.Println("the sum of all values in inputs slice =>", sum(inputs...))

	fmt.Println("all even numbers in inputs slice =>", even(inputs...))
	fmt.Println("the sum of even numbers in inputs slice =>", sum(even(inputs...)...))

	fmt.Println("the sum of all even values in input slice, using sumFn() =>", sumFn(even, inputs...))
	fmt.Println("the sum of all odd values in input slice, using sumFn() =>", sumFn(odd, inputs...))
}

func sum(s ...int) int {
	result := 0
	for _, val := range s {
		result += val
	}
	return result
}

func even(n ...int) []int {
	var result []int
	for _, val := range n {
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

func odd(n ...int) []int {
	var result []int
	for _, val := range n {
		if val%2 != 0 {
			result = append(result, val)
		}
	}
	return result
}

func sumFn(fn func(n ...int) []int, xi ...int) int {
	return sum(fn(xi...)...)
}

/*
func sum2(<fn>, <slice>) {...}

e.g.
func sum2(even, []int{1, 2, 3, 4})
=> 6

func sum2(odd, []int{1, 2, 3, 4})
=> 4
*/
