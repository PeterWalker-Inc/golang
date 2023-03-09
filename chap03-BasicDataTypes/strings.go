package main

import "fmt"

func main() {
	var s string = "Hello, World"
	fmt.Println(s[:5], s[7:])
	fmt.Println("Alert is \\a \a, backspace is \\b s\b, form feed is \\f \f, tab is \\t \t, vertical tab is \\v \v")

	// raw string
	fmt.Println(`Alert is \\a \a, backspace is \\b s\b, form feed is \\f \f, tab is \\t \t, vertical tab is \\v \v`)

	fmt.Println(HasPrefix(s, "Hello"))
	fmt.Println(HasSuffix(s, "World"))
	fmt.Println(Contains(s, "orl"))

	// Hex
	fmt.Printf("Hex of பீட்டர்: % x\n", "பீட்டர்")
	fmt.Println(string(65))
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
