package main

import "fmt"

func main() {
	var s string
	if s != "" && s[0] == 'x' { //False, panic if s != ""
		fmt.Println()
	}
	fmt.Println(s+"is")

	b := "not empty string"
	if b != "" {
		fmt.Println(b)
	}

	// There is no implicit converison of non-zero interger to boolean
	if itob(1) {
		fmt.Println(btoi(itob(1)))
	}

}

// Boolean to Integer
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Integer to boolean
func itob(i int) bool {
	return i != 0
}