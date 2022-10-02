package main

import "fmt"

func main() {
	num := 42
	name := "golang"
	xi := []int{1, 2, 3}

	printAllDetails(num)
	printAllDetails(&num)

	printAllDetails(name)
	printAllDetails(&name)

	printAllDetails(xi)
	printAllDetails(&xi)
}

func sGetType(n interface{}) string {
	return fmt.Sprintf("%T", n)
}

func printType(n interface{}) {
	t := sGetType(n)

	switch t {
	case "int":
		fmt.Println(n, "is an int")
	case "*int":
		fmt.Println(n, "is a pointer of type *int")
	case "string":
		fmt.Println(n, "is a string")
	case "*string":
		fmt.Println(n, "is a pointer of type *string")
	case "[]int":
		fmt.Println(n, "is slice []int")
	case "*[]int":
		fmt.Println(n, "is a pointer of type *[]int")
	}
}

func printAllDetails(n interface{}) {
	printType(n)
	t := sGetType(n)
	switch t {
	case "int":
		fmt.Println(n, "is an int")
		fmt.Printf(" => value: %d\n", n)
	case "string":
		fmt.Printf(" => value: %s\n", n)
	default:
		fmt.Printf(" => value: %v\n", n)
	}
	fmt.Printf(" => memory location: %v\n\n", &n)
}
