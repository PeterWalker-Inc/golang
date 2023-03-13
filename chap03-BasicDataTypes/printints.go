package main 

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(printints([]int{1,2,3}))
}

func printints(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range(values) {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
