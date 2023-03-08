package main

import "fmt"

func main() {
	var f float64 = 167772122323
	fmt.Printf("f is %f\n", f)
	fmt.Printf("f is %e\n", f)
	fmt.Printf("f is %g\n", f)

	var n float32 = 99999999 // upto 6 digit exact precision
	fmt.Println(n == n+1)
	var m float64 = 9999999999999999 // upto 15 digit exact precision
	fmt.Println(m == m+1)
}