package main

import "fmt"


func s() *string{
	var v string
	fmt.Println(&v)
	return &v
}

func i() *int64{
	var i int64
	fmt.Println(&i)
	return &i
}

func incr(n *int) {
	*n++
}

func main() {
	value := 4;
	p := &value
	*p = 5
	fmt.Println(value)
	fmt.Println(s() == s())
	fmt.Println(i() == i())

	n := 5
	incr(&n)
	fmt.Printf("n is %d\n", n)
}