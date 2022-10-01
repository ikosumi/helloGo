package main

import "fmt"

func main() {
	inputs := []int{1, 2, 3, 4}
	fmt.Println("the inputs slice is", inputs)
	fmt.Println("the sum of all values in inputs slice is", sum(inputs...))

	fmt.Println("all even numbers in inputs slice are", even(inputs...))
	fmt.Println("the sum of even numbers in inputs slice is", sum(even(inputs...)...))
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
