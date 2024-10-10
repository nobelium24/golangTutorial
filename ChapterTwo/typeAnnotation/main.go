package main

import "fmt"

// Once a type is created as seen below, it automatically has a conversion operation T(x) that can convert a value (with the same underlying type as the type or pointers to the same underlying type) to itself
type Celsius float64
type Fahrenheit float64

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Methods can be associated with a defined type and can then be used on an instance of the defined type
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	c := Celsius(45.3)
	fmt.Print(c.String())
}
