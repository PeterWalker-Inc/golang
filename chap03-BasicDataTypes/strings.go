package main

import "fmt"

func main() {
	var s string = "Hello, World"
	fmt.Println(s[:5], s[7:])
	fmt.Println("Alert is \\a \a, backspace is \\b s\b, form feed is \\f \f, tab is \\t \t, vertical tab is \\v \v")

	// raw string
	fmt.Println(`Alert is \\a \a, backspace is \\b s\b, form feed is \\f \f, tab is \\t \t, vertical tab is \\v \v`)

	const GoBook = ` The Go Programming language
	Alan and Kerninghan
	`
	fmt.Println(GoBook)
}