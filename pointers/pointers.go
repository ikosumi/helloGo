package main

import "fmt"

func main() {
	n := 42          // assigns value to a
	addressOfn := &n // saves memory address of a

	fmt.Println("current value of a:", n)
	fmt.Println("address of a using addressOfn: ", addressOfn)
	fmt.Println("de-referenced value of addressOfn:", *addressOfn)

	fmt.Println("setting value of a using it's address memory...")
	*addressOfn = 43           // updates the stored value of addressOfn
	fmt.Println("a is now", n) // a is now 43

}
