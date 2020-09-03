package lenconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

func MToFt(m Meter) Feet  { return Feet(m / 0.3048) }
func FtToM(ft Feet) Meter { return Meter(ft * 0.3048) }
