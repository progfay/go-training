package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

var (
	AbsoluteZeroF Fahrenheit = CToF(AbsoluteZeroC)
	FreezingF     Fahrenheit = CToF(FreezingC)
	BoilingF      Fahrenheit = CToF(BoilingC)
)

var (
	AbsoluteZeroK Kelvin = CToK(AbsoluteZeroC)
	FreezingK     Kelvin = CToK(FreezingC)
	BoilingK      Kelvin = CToK(BoilingC)
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
