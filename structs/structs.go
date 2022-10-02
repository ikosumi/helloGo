package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type Employee struct {
	id     int
	person Person
}

func main() {
	ilir := Employee{1, Person{"Ilir", "Kosumi"}}
	fmt.Println(ilir)
}
