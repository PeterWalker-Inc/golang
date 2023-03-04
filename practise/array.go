package main

import (
	"fmt"
	"sort"
)

func main() {
	var iSlice = []int{124, 35, 2456, 652, 35}
	iSlice = append(iSlice, 45)
	fmt.Println(iSlice)

	var stringSlice = []string{"Red", "Blue", "Green"}
	stringSlice = append(stringSlice, "Emerald")
	stringSlice = append(stringSlice[1:])
	fmt.Println(stringSlice)

	numbers := make([]int, 5, 9)
	numbers[0] = 7
	numbers[1] = 6
	numbers[2] = 7
	numbers[3] = 6
	numbers[4] = 7

	numbers = append(numbers, 34)
	numbers = append(numbers, 35)
	numbers = append(numbers, 36)
	numbers = append(numbers, 37)
	numbers = append(numbers, 37)
	numbers = append(numbers, 37)
	numbers = append(numbers, 37)
	numbers = append(numbers, 37)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
