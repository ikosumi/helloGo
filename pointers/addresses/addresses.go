package main

import "fmt"

func main() {
	num := 1

	fmt.Printf("current value of num %d in address %v\n", num, &num)
	UpdateValueInAddr(&num, 2)
	fmt.Println("new value of num", num)

}

func UpdateValueInAddr(address *int, value int) {
	fmt.Printf("writing value %d into address %v\n", value, address)
	*address = value
}
