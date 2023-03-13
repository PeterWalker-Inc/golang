package main

import (
	"fmt"
	"strings"
)
func main() {
	path := "/home/peter/rep/golang/chap03-BasicDataTypes/basename2.go"
	fmt.Println(basename(path))
}
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
