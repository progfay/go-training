package complex

import "math/big"

type BigFloatComplex struct {
	real, imagine *big.Float
}

func NewBigFloatComplex(real, imagine float64) *BigFloatComplex {
	return &BigFloatComplex{
		real:    big.NewFloat(real),
		imagine: big.NewFloat(imagine),
	}
}

func (c1 *BigFloatComplex) Add(c2 *BigFloatComplex) *BigFloatComplex {
	ret := NewBigFloatComplex(0, 0)
	ret.real.Add(c1.real, c2.real)
	ret.imagine.Add(c1.imagine, c2.imagine)

	return ret
}

func (c1 *BigFloatComplex) Mul(c2 *BigFloatComplex) *BigFloatComplex {
	ret := NewBigFloatComplex(0, 0)
	ret.real.Sub(
		big.NewFloat(0).Mul(c1.real, c2.real),
		big.NewFloat(0).Mul(c1.imagine, c2.imagine),
	)
	ret.imagine.Add(
		big.NewFloat(0).Mul(c1.real, c2.imagine),
		big.NewFloat(0).Mul(c1.imagine, c2.real),
	)

	return ret
}

func (c *BigFloatComplex) Abs() float64 {
	var s, t = big.NewFloat(0), big.NewFloat(0)
	s.Abs(c.real)
	t.Abs(c.imagine)
	s.Mul(s, s)
	t.Mul(t, t)
	s.Add(s, t)
	ret, _ := s.Sqrt(s).Float64()

	return ret
}
