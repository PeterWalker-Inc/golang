package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	// Write first group of numbers
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	// Write the rest
	for j := i + 3; j < n; {
		buf.WriteString("," + s[i:j])
		i, j = j, j+3
	}
	buf.WriteString("," + s[i:])
	return buf.String()
}

func main() {
	num := "12345678904124254235345234"
	fmt.Println(num)
	fmt.Println(comma(num))
}
