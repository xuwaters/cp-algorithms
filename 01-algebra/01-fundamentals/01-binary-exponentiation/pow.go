package exp

func Pow(b Powable, n int) Powable {
	exp := b
	res := b.One()
	for n > 0 {
		if (n & 1) == 1 {
			res = res.Multiply(exp)
		}
		exp = exp.Multiply(exp)
		n >>= 1
	}
	return res
}

type Powable interface {
	Multiply(other Powable) Powable
	One() Powable
}

type PowableInt int

var _ Powable = (PowableInt)(0)

func (i PowableInt) One() Powable                   { return PowableInt(1) }
func (i PowableInt) Multiply(other Powable) Powable { return i * other.(PowableInt) }
