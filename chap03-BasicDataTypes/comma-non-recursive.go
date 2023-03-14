package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer
	wordcount := 1
	for _, v := range s {
		wordcount++
		if wordcount % 3 == 0 {
			buf.WriteRune(',') // character writing
		}
		buf.WriteRune(v) // character writing
	}
	return buf.String()
}

func main() {
	num := "1234567890"
	fmt.Println(comma(num))
}
