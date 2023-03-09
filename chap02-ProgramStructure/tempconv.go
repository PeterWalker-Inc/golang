package main

import "fmt"

type Celsius float64
type Farenheit float64
type age int

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g°C\n", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%g°F\n", f) }

func CToF(c Celsius) Farenheit { return Farenheit((c * 9 / 5) + 32) }
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Println(CToF(100))
}
