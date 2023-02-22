package main

import "fmt"
import "time"

func main() {
	n := time.Now()
	fmt.Println("Now it is: ", n)
	t := time.Date(2023, time.February, 21, 3, 10, 15, 0, time.UTC)
	fmt.Println("Today is: ", t);
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Feb 21 03:10:15 2023")
	fmt.Println("Parsed Time is: ", parsedTime)
	fmt.Printf("Type of parsed time is: %T", parsedTime)

}