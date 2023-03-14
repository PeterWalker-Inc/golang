package main

import "fmt"
import "time"

func main() {
	const (
		a = 1
		b
		c
		d = 2
		e
		f
	)
	fmt.Println(a,b,c,d,e,f)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("no Delay is %T %[1]v\n", noDelay)
	fmt.Printf("timeout is %T %[1]v\n", timeout)
	fmt.Printf("time.Minute is %T %[1]v\n", time.Minute)

}
