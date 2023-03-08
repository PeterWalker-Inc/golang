package main

import "fmt"

func main() {
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2

	fmt.Printf("x is %08b\n", x)
	fmt.Printf("y is %08b\n", y)
	fmt.Printf("^y is %08b\n", ^y)
	fmt.Println("---------------")
	fmt.Printf("x&y (intersection) %08b\n", x&y)
	fmt.Printf("x|y (union) %08b\n", x|y)
	fmt.Printf("x^y (symmetric difference) %08b\n", x^y)
	fmt.Printf("x&^y (difference) %08b\n", x&^y)
}