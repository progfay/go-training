package complex

import (
	"math"
	"math/big"
)

type BigRatComplex struct {
	real, imagine *big.Rat
}

func NewBigRatComplex(real, imagine float64) *BigRatComplex {
	return &BigRatComplex{
		real:    (&big.Rat{}).SetFloat64(0),
		imagine: (&big.Rat{}).SetFloat64(0),
	}
}

func (c1 *BigRatComplex) Add(c2 *BigRatComplex) *BigRatComplex {
	ret := NewBigRatComplex(0, 0)
	ret.real.Add(c1.real, c2.real)
	ret.imagine.Add(c1.imagine, c2.imagine)

	return ret
}

func (c1 *BigRatComplex) Mul(c2 *BigRatComplex) *BigRatComplex {
	ret := NewBigRatComplex(0, 0)
	ret.real.Sub(
		(&big.Rat{}).SetFloat64(0).Mul(c1.real, c2.real),
		(&big.Rat{}).SetFloat64(0).Mul(c1.imagine, c2.imagine),
	)
	ret.imagine.Add(
		(&big.Rat{}).SetFloat64(0).Mul(c1.real, c2.imagine),
		(&big.Rat{}).SetFloat64(0).Mul(c1.imagine, c2.real),
	)

	return ret
}

func (c *BigRatComplex) Abs() float64 {
	s, _ := c.real.Float64()
	t, _ := c.imagine.Float64()
	return math.Hypot(s, t)
}
