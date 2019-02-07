package fibonacci

// F(0) = 0
// F(1) = 1
// F(n+1) = F(n) + F(n-1)
//
//  [[Fn-1, Fn],
//   [Fn,  Fn+1]]  = M^n
//
//  where M = [[0, 1],
//             [1, 1]]
//
//  [Fn, Fn+1] = [0 1] * M^n
//

// copy from 01-binary-exponentiation

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	mat := Mat2x2{
		{1, 1},
		{1, 0},
	}
	mat = Pow(mat, n-1).(Mat2x2)
	return mat[0][0]
}

type Mat2x2 [2][2]int

var _ Powable = (*Mat2x2)(nil)

func (m Mat2x2) One() Powable {
	return Mat2x2{
		{1, 0},
		{0, 1},
	}
}

func (m Mat2x2) Multiply(o Powable) Powable {
	m2 := o.(Mat2x2)
	return Mat2x2{
		{m[0][0]*m2[0][0] + m[0][1]*m2[1][0], m[0][0]*m2[0][1] + m[0][1]*m2[1][1]},
		{m[1][0]*m2[0][0] + m[1][1]*m2[1][0], m[1][0]*m2[0][1] + m[1][1]*m2[1][1]},
	}
}

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
