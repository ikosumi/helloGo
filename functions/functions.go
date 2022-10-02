package main

import "fmt"

func main() {
	inputs := []int{1, 2, 3, 4}
	fmt.Println("the inputs slice is", inputs)
	fmt.Println("the sum of all values in inputs slice =>", Sum(inputs...))

	fmt.Println("all even numbers in inputs slice =>", Even(inputs...))
	fmt.Println("the sum of even numbers in inputs slice =>", Sum(Even(inputs...)...))

	fmt.Println("the sum of all even values in input slice, using sumFn() =>", SumFn(Even, inputs...))
	fmt.Println("the sum of all odd values in input slice, using sumFn() =>", SumFn(Odd, inputs...))
}

func Sum(s ...int) int {
	result := 0
	for _, val := range s {
		result += val
	}
	return result
}

func Even(n ...int) []int {
	var result []int
	for _, val := range n {
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

func Odd(n ...int) []int {
	var result []int
	for _, val := range n {
		if val%2 != 0 {
			result = append(result, val)
		}
	}
	return result
}

func SumFn(fn func(n ...int) []int, xi ...int) int {
	return Sum(fn(xi...)...)
}
