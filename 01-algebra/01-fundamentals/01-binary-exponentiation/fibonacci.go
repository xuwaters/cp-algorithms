package exp

// f(0) = 1
// f(1) = 1
// f(2) = 2
// f(3) = 3
// f(4) = 5
// f(n) = f(n-1) + f(n-2)
//
// [1, 0]
// [fn, fn_1] = [ fn_1, fn_2 ] * [ [1, 1], [1, 0] ]
//
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
