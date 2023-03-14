package main

import (
	"fmt"
	"strings"
)

func hasSign(s string) bool {
	sign := strings.Index(s, "-")
	if sign != 0{
		return false
	}
	return true
}

func dotIndex(s string) int {
	dot := strings.Index(s, ".")
	return dot
}

func comma(s string) string{
	dot := dotIndex(s)
	var decimal string
	if dot >= 0 {
		decimal = s[dot:]
		s = s[:dot]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] + decimal
}

func main() {
	s := "-12345678.90"
	signed := hasSign(s)
	if signed {
		s = s[1: ]
		s = comma(s)
		s = "-" + s
	}else {
		s = comma(s)
	}
	fmt.Println(s)
}
