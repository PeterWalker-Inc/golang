package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	hello := "Hello"
	fmt.Printf("%v", hello)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter num 1: ")
	input1, _ := reader.ReadString('\n')
	val1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		panic(err)
	}
	fmt.Print("Enter num 2: ")
	input2, _ := reader.ReadString('\n')
	val2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic(err)
	}
	sum := float64(val1) + float64(val2)
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum is %v", sum)
}